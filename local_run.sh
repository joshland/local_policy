#!/bin/bash

source activate

sudo $(which ansible-playbook) --connection=local --inventory 127.0.0.1, local_check.yml --extra-vars "@vars.yaml"  $* --diff 
