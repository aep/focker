package main


import (
	"os"
	"path/filepath"
	"fmt"
	"crypto/sha256"
	"os/exec"
	"io/ioutil"
)



func dockerHasImage(name string) bool {
	cmd := exec.Command("docker", "images", "-q", name)
	out, err := cmd.Output()
	if err != nil {
		return false
	}
	return len(out) > 0
}


func findFockerfile(pwd string) (string, []byte, error) {

	content, err := ioutil.ReadFile(pwd + "/Focker/Dockerfile")
	if err == nil {
		return pwd, content, nil
	}

	if pwd == "/" {
		return "", nil, fmt.Errorf("Focker not found")
	}

	return findFockerfile(filepath.Dir(pwd))
}


func main() {
	pwd , err := os.Getwd()
	if err != nil {
		panic(err);
	}

	nwd, content, err := findFockerfile(pwd)
	if err != nil {
		fmt.Println(err);
		os.Exit(1)
	}

	hasher := sha256.New()
	hasher.Write(content)
	hash := fmt.Sprintf("%x", hasher.Sum(nil))


	fmt.Println("using builder focker-" + hash + " from ", nwd + "/Focker/Dockerfile")


	if !dockerHasImage("focker-" + hash) {
		os.Chdir(nwd + "/Focker")

		dockerfile, err := ioutil.TempFile("", "Dockerfile")
		if err != nil {
			panic(err)
		}
		defer os.Remove(dockerfile.Name())

		dockerfile.Write([]byte(`
`))



		dockerfile.Write(content)

		dockerfile.Write([]byte(`
run apt update && apt install -y sudo
run useradd --home-dir /src -m --uid 1000 user 
workdir /src
run echo "user ALL=(ALL) NOPASSWD: ALL" >> /etc/sudoers
user user
		`));
		dockerfile.Close()


		cmd := exec.Command("docker", "build", "-t", "focker-" + hash, "-f", dockerfile.Name(), ".")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			panic(err)
		}
	}

	cmd := exec.Command("docker", "run", "-it", "--rm", "-v", pwd + ":/src", "focker-" + hash)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Run()
	os.Exit(cmd.ProcessState.ExitCode())

}
