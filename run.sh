#!/bin/bash

ANSIBLE_PYTHON_INTERPRETER=/usr/bin/python3

sudo $(which ansible-playbook) --connection=local --inventory 127.0.0.1, local_check.yml $*
