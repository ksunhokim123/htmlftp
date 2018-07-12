import click
import os
import sys
import json
from threading import Thread

from mouseftp.ftp import ftp_server
import mouseftp.api

def read_dir(ctx, param, value):
    if not value:
        if not os.path.exists('home'):
            os.mkdir('home')
        if not os.path.isdir('home'):
            print('home is not a dir')
            sys.exit(1)
        value = 'home'
    return value

@click.command()
@click.option('-d', '--dir', callback=read_dir, type=click.Path(exists=True, file_okay=False),
    help='user home folder path')
@click.option('-fi', '--ftpaddress', type=(str, int), default=('localhost', 21),
    help='ftp (address, port)')
@click.option('-ai', '--apiaddress', type=(str, int), default=('localhost', 80),
    help='api (address, port)')
def main(**kwargs):
    server = ftp_server(kwargs['ftpaddress'], kwargs['dir'])
    thread = Thread(target=server.run)
    thread.start()

    mouseftp.api.start(server, kwargs['apiaddress'])
