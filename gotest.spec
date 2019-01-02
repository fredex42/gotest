%define name gotest
%define version 1.0
%define unmangled_version 1.0
%define release 2

Summary: Test program for Go with packaging
Name: %{name}
Version: %{version}
Release: %{release}
License: Internal GNM software
Group: Applications/Productivity
BuildRoot: %{_tmppath}/gnmawsgr
Prefix: %{_prefix}
BuildArch: x86_64
Vendor: Andy Gallagher <andy.gallagher@theguardian.com>

%description
Test commandline program in Go

%prep
rm -rf $RPM_BUILD_DIR/gotest

%build
cp -a $RPM_SOURCE_DIR/gotest $RPM_BUILD_DIR/gotest
cd $RPM_BUILD_DIR/gotest
go build .

%install
mkdir -p $RPM_BUILD_ROOT/usr/local/bin
mv $RPM_BUILD_DIR/gotest/gotest $RPM_BUILD_ROOT/usr/local/bin

%clean
rm -rf $RPM_BUILD_ROOT

%files
%defattr(-,root,root)
/usr/local/bin/gotest

%pre

%post

%preun
