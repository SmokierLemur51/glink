import subprocess
from pathlib import Path

def new_poetry_project() -> None:
    """
    """
    # getting projects directory
    home = Path.home()
    project_directory = home / 'glink'
    # name new projects
    project = input("Name of new project: ")
    poetry_package = input("Name of poetry package: ")
    # create path to new project dir & create if not exists
    new_project_dir = project_directory / project
    new_project_dir.mkdir(parents=True, exist_ok=True)
    # path to new poetry package in new_project_dir
    new_package = new_project_dir / poetry_package

    try:
        result = subprocess.run(
            ['poetry', 'new', '--src', new_package],
            check=True,
            text=True,
            capture_output=True,    
        )
        print(f"Successfully created new project [{project}] with poetry package [{poetry_package}]")
        print(f"Output: {result.stdout}")
    except subprocess.CalledProcessError as e: 
        print(f"Failed with error message:\n{e}")
        print(f"Eror output: {e.stderr}")


def install_requirements() -> None:
    with open('requirements.txt', 'r') as f:
        for x in f:
            subprocess.run(['pip', 'uninstall', x.strip()], check=True)


new_poetry_project()
