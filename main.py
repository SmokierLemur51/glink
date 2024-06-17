import subprocess

# with open("output/x.py", "w") as x:
#     data = None
#     with open("templates/python/flask/manage.py", "r") as d:
#         data = d.read()
#     x.write(data)

package = input("Name Poetry package: ")
subprocess.run([
    "cd output/", 
    "&&", 
    "poetry new --src {}".format(package), 
    "&&", 
    "cd {}/src/{}".format(package, package),
])

