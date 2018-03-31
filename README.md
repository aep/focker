from "fucker", plus docker

i use these to build shitty software that requires ubuntu or is generally unmaintained.
the docker images have a volume on /user which will be mounted to host pwd.
It used to be different images for different use cases,
but now its a giant kitchen sink of shit. whatever.



This isn't exactly intented to be used by other people, but if you insist:


installing:
```
make
```


using:

```
$ cd anyproject
$ ./configure
error: we detected you're not using ubuntu 14.04 and will therefor refuse to cooperate
$ fock ./configure
praised be our lord shuttleworth.
```
