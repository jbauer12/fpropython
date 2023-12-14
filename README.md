Python Distribution is needed. 
Unfortunately Docker cant be (easily) used because of the GUI. 
Therefore we will use virtual environments.

pip install virtualenv
python3 -m venv .venv
source .venv/bin/activate
pip install -r requirements.txt
