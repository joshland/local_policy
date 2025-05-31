# Local Base Checks

A quick ansible playbook which updates the local host to set some basic security requirements

## Execution

    sudo $(which ansible-playbook) --connection=local --inventory 127.0.0.1, local_check.yml --extra-vars "@vars.yaml"  $* --diff --check
