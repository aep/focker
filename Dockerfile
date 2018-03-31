FROM    ubuntu:14.04

RUN dpkg --add-architecture i386 && \
    echo "foreign-architecture i386" > /etc/dpkg/dpkg.cfg.d/multiarch && \
    apt-get update && apt-get install -y \
        build-essential \
        sudo \
        python \
        software-properties-common \
        python-software-properties \
        bison flex unzip zip gperf libxml2-utils \
        libc6:i386 libncurses5:i386 libstdc++6:i386 zlib1g:i386

RUN mkdir -p /var/cache/oracle-jdk6-installer/
ADD jdk-6u45-linux-x64.bin /var/cache/oracle-jdk6-installer/

RUN echo debconf shared/accepted-oracle-license-v1-1 select true | sudo debconf-set-selections
RUN echo debconf shared/accepted-oracle-license-v1-1 seen true | sudo debconf-set-selections

RUN     sudo add-apt-repository ppa:webupd8team/java && \
        sudo apt-get update && \
        sudo apt-get install  -y oracle-java6-installer && \
        sudo apt-get install  -y oracle-java6-set-default


RUN     echo 'user ALL=(ALL) NOPASSWD: ALL' > /etc/sudoers
RUN     useradd user -u 1000 -U -m -d /user
VOLUME  /user
WORKDIR /user
USER    user
ENV     USER=user
