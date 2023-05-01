Sure, here's a basic outline of how the Python classes and methods might look for the "Virtualenv Health Checker" application. Please note that the actual implementation would require handling more details and edge cases than are represented here.

```python
import os
from typing import List, Dict

class VirtualEnvironment:
    """
    Class representing a Python virtual environment.
    """
    def __init__(self, path: str):
        """
        Initialize a virtual environment.

        :param path: The path to the virtual environment.
        """
        self.path = path

    def get_installed_packages(self) -> Dict[str, str]:
        """
        Get a dictionary of all packages installed in the virtual environment,
        with package names as keys and versions as values.

        :return: A dictionary of packages and their versions.
        """
        pass  # implementation goes here

class VirtualenvHealthChecker:
    """
    Class that checks the health of Python virtual environments.
    """
    def __init__(self, envs: List[VirtualEnvironment]):
        """
        Initialize the health checker with a list of virtual environments.

        :param envs: A list of VirtualEnvironment objects to check.
        """
        self.envs = envs

    def check_for_outdated_packages(self) -> Dict[str, List[str]]:
        """
        Check for outdated packages in all virtual environments.

        :return: A dictionary where keys are environment paths and values are
                 lists of outdated packages in those environments.
        """
        pass  # implementation goes here

    def check_for_security_vulnerabilities(self) -> Dict[str, List[str]]:
        """
        Check for packages with known security vulnerabilities in all virtual environments.

        :return: A dictionary where keys are environment paths and values are
                 lists of packages with security vulnerabilities in those environments.
        """
        pass  # implementation goes here

    def check_for_unused_dependencies(self) -> Dict[str, List[str]]:
        """
        Check for unused dependencies in all virtual environments.

        :return: A dictionary where keys are environment paths and values are
                 lists of unused dependencies in those environments.
        """
        pass  # implementation goes here
```

This code provides the structure for the application, but the actual implementation of each method would require additional code. For example, the `get_installed_packages` method might use the `pip` command to list all installed packages, the `check_for_outdated_packages` method might use the `pip list --outdated` command to find outdated packages, etc. Security vulnerabilities can be checked using various databases like Safety DB, or through APIs of services like PyUp's Safety, and unused dependencies would need some form of static code analysis to determine if imported modules are actually used.