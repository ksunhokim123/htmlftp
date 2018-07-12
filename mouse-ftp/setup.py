from setuptools import setup

setup(
  name = 'mouseftp',
  packages = ['mouseftp'], # this must be the same as the name above
  version = '0.2',
  install_requires=[
        'Click',
        'pyftpdlib',
        'flask'
    ],
  entry_points = {
    "console_scripts": ['mouseftp = mouseftp.main:main']
    },
  description = 'a ftp server which provides user api through http',
  author = 'Sunho Kim',
  author_email = 'ksunhokim123@gmail.com',
  url = 'https://github.com/sunho/mouse-ftp', # use the URL to the github repo
  download_url = 'https://github.com/sunho/mouse-ftp/archive/0.3.tar.gz', # I'll explain this in a second
  keywords = [], # arbitrary keywords
  classifiers = [],
)
