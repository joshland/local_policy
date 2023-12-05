#!/bin/bash

[ -e ".venv" ] && rm -fR .venv/

python3 -m venv .venv

[ ! -e "activate" ] && ln -s .venv/bin/activate

source activate

pip install --upgrade pip wheel

[ -e 'requirements.txt' ] && pip  install --upgrade -r requirements.txt
