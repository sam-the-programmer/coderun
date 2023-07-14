import os
import shutil
import subprocess


def ask():
	is_run = input("\033[35m Clearing old binaries. OK to continue? (Y/n) \033[0m")

	if is_run.lower() == "n":
		exit()
	elif is_run.lower() in ["y", " ", ""]:
		return 0
	else:
		ask()


def build_apps():
	# windows
	print("\033[33m Building Windows Executable \033[0m")
	os.environ["GOOS"] = "windows"
	subprocess.run(["go", "build", "coderun.go"])

	# macos
	print("\033[33m Building MacOS Executable \033[0m")
	os.environ["GOOS"] = "darwin"
	subprocess.run(["go", "build", "coderun.go"])
	os.rename(
		"coderun", "coderun-macos"
	)  # to avoid confusion with the linux executable

	# linux
	print("\033[33m Building Linux Executable \033[0m")
	os.environ["GOOS"] = "linux"
	subprocess.run(["go", "build", "coderun.go"])


def main():
	ask()
	for file in os.listdir("bin"):
		try:
			os.remove(f"bin/{file}")
		except PermissionError:  # if it is protected or a folder or something
			pass

	print("\033[31m Outdated binaries cleared. Preparing to build new files. \033[0m")
	build_apps()

	print("\033[36m Fixing file tree... \033[0m")
	shutil.copy("coderun", "bin/")
	shutil.copy("coderun-macos", "bin/")
	shutil.copy("coderun.exe", "bin/")

	os.remove("coderun")
	os.remove("coderun-macos")
	os.remove("coderun.exe")

	print("\033[32m Binaries built and sorted! \033[0m")


if __name__ == "__main__":
	main()
