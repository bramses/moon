# [virtualenv](https://click.palletsprojects.com/en/8.1.x/quickstart/#virtualenv)

Virtualenv is probably what you want to use for developing Click applications.

What problem does virtualenv solve? Chances are that you want to use it for other projects besides your Click script. But the more projects you have, the more likely it is that you will be working with different versions of Python itself, or at least different versions of Python libraries. Let’s face it: quite often libraries break backwards compatibility, and it’s unlikely that any serious application will have zero dependencies. So what do you do if two or more of your projects have conflicting dependencies?

Virtualenv to the rescue! Virtualenv enables multiple side-by-side installations of Python, one for each project. It doesn’t actually install separate copies of Python, but it does provide a clever way to keep different project environments isolated. Let’s see how virtualenv works.