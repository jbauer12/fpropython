# Checkers Game with Minimax Algorithm

This code implements the board game checkers with the minimax algorithm in Python

## Installation
After the Installation of Python (used 3.10) [Python](https://www.python.org/downloads/)

Use the package manager [pip](https://pip.pypa.io/en/stable/) to install [virtualenv](https://docs.python.org/3/library/venv.html) (the recommend way - unfortunately Docker can not be used easily through the usage of GUI).
virtualenv

The documentation describes virtualenv as following: "A virtual environment is created on top of an existing Python installation, known as the virtual environment’s “base” Python, and may optionally be isolated from the packages in the base environment, so only those explicitly installed in the virtual environment are available."

It helps us separating our dependencies for a specific project from global ones. 

```bash
pip install virtualenv
```
After that instantiate and activate the new virtual environment.
```bash
#If not already installed
pip install virtualenv

# Instantiate virtualenv called .venv
cd Python
python3 -m venv .venv

#Activate the virtualenv

#If you are using Linux distribution like Ubuntu or MacOS
source .venv/bin/activate

#If you are using Windows
.\.venv\Scripts\activate

# Install all requirements for running the code 
pip install -r requirements.txt
cd ..

```


## Usage Of Code

```bash
#from the root directory where you can finde the folders Python and Go
#Run the game
python Python/src/main.py

```
## Running Tests

```bash
#from the root directory where you can finde the folders Python and Go

#Run the game for Playing
python Python/src/main.py

#Run Tests
python Python/src/test.py

```

