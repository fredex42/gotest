FROM centos:7

RUN yum -y update && yum -y groupinstall 'Development Tools' && yum -y install wget shasum && yum -y clean all && rm -rf /var/cache/yum
COPY install_go.sh /usr/local/bin
RUN chmod a+x /usr/local/bin/install_go.sh; /usr/local/bin/install_go.sh
RUN adduser build
USER build
RUN echo 'declare -x PATH=$PATH:/usr/local/go/bin' >> /home/build/.bash_profile && mkdir -p /home/build/rpmbuild/SOURCES
USER root
ENTRYPOINT /bin/bash --login
