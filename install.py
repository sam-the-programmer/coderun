import shutil
import platform

if platform.system() == "Windows":
	shutil.copy("bin/coderun.exe", r"%localappdata%/Programs/Python/Python39/Scripts")
elif platform.system() == "Linux":
	shutil.copy("bin/coderun", "/usr/bin")
elif platform.system() == "Java":
	shutil.copy("bin/coderun-macos", "/usr/bin")