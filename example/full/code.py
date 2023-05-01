import os
import unittest
from typing import List, Dict
import subprocess

class VirtualEnvironment:
    def __init__(self, path: str):
        """
        Initialize a virtual environment.

        :param path: The path to the virtual environment.
        """
        if not os.path.exists(path):
            raise ValueError(f"The provided path '{path}' does not exist.")
        self.path = path


    def get_installed_packages(self) -> Dict[str, str]:
        """
        Get a dictionary of all packages installed in the virtual environment,
        with package names as keys and versions as values.

        :return: A dictionary of packages and their versions.
        """
        # Implementation here would depend on how you're getting the installed packages.
        # For now, I'm simulating it with a subprocess call to pip freeze.
        try:
            result = subprocess.check_output(['pip', 'freeze'], cwd=self.path)
            packages = dict(line.split('==') for line in result.decode().splitlines())
            return packages
        except subprocess.CalledProcessError as e:
            raise RuntimeError(f"Error retrieving installed packages: {e}")



class TestVirtualEnvironment(unittest.TestCase):
    def test_init_invalid_path(self):
        with self.assertRaises(ValueError):
            VirtualEnvironment("/path/that/does/not/exist")

    def test_get_installed_packages(self):
        env = VirtualEnvironment("/valid/path")  # Replace with a valid path
        packages = env.get_installed_packages()
        self.assertIsInstance(packages, dict)
        for key, value in packages.items():
            self.assertIsInstance(key, str)
            self.assertIsInstance(value, str)

# And so on for the other tests...

if __name__ == "__main__":
    unittest.main()