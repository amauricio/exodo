#!/usr/bin/env bash

version='lsb_release --release | cut -f2'
codename='lsb_release --codename | cut -f2'
deb http://apt.postgresql.org/pub/repos/apt/ $version-pgdg main