import shutil
import platform

if platform.system() == "Windows":
	shutil.copyfile("bin/coderun.exe", r"%localappdata%/Programs/Python/Python39/Scripts")
elif platform.system() == "Linux":
	shutil.copyfile("bin/coderun", "/usr/bin")
elif platform.system() == "Java":
	shutil.copyfile("bin/coderun-macos", "/usr/bin")