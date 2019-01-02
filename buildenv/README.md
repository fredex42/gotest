# GoTest build environment

This directory contains a simple build environment that is suitable for building Go apps into RPMs.

## How to use it

- Build the image or locate it on Docker Hub. `{image-id}` will be of the form org/image:version
- Run `docker run --rm -it -v {path-to-your-sources}:/mnt {image-id}` to get a shell
- You're logged in as the build user.
- Copy your sources to the RPM SOURCES directory: `mkdir ~/rpmbuild/SOURCES/gotest; cp -a /mnt ~/rpmbuild/SOURCES/gotest`
- Go to the sources and run rpmbuild: `cd /mnt; rpmbuild -bb gotest.spec`.
- This should output the final built RPM into `~/rpmbuild/RPMS`
- Consult the details of `gotest.spec` to see how the build is done, what source directory is used etc.

## How to build it

- Nothing special is required to build, other than a working Docker installation.
- Simply cd into this directroy (`buildenv`) and run `docker build . -t yourname/gotest-build:1`.  An up-to-date centos image should be downloaded and the necessary stuff installed.
- You can then run `docker run --rm -it -v {path-to-your-sources}:/mnt yourname/gotest-build:1` in the step 2 command above
