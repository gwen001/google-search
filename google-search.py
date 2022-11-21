#!/usr/bin/python3

import os
import sys
import json
import argparse
import urllib.parse
from goop import goop
from functools import partial
from multiprocessing.dummy import Pool
from colored import fg, bg, attr


# def banner():
# 	print("""
#                          _                                _
#   __ _  ___   ___   __ _| | ___   ___  ___  __ _ _ __ ___| |__        _ __  _   _
#  / _` |/ _ \ / _ \ / _` | |/ _ \ / __|/ _ \/ _` | '__/ __| '_ \      | '_ \| | | |
# | (_| | (_) | (_) | (_| | |  __/ \__ \  __/ (_| | | | (__| | | |  _  | |_) | |_| |
#  \__, |\___/ \___/ \__, |_|\___| |___/\___|\__,_|_|  \___|_| |_| (_) | .__/ \__, |
#  |___/             |___/                                             |_|    |___/

#                         by @gwendallecoguic

# """)
# 	pass


parser = argparse.ArgumentParser()
# parser.add_argument( "-r","--raw",help="remove banner", action="store_true" )
parser.add_argument( "-s","--search",help="search term" )
parser.add_argument( "-d","--decode",help="urldecode the results", action="store_true" )
parser.add_argument( "-c","--fbcookie",help="your facebook cookie" )

parser.parse_args()
args = parser.parse_args()

# if not args.raw:
#     banner()

if args.fbcookie:
    fb_cookie = args.fbcookie
else:
    fb_cookie = os.getenv('FACEBOOK_COOKIE')
if not fb_cookie:
    parser.error( 'facebook cookie is missing' )

if args.search:
    search = args.search
else:
    parser.error( 'search expression is missing' )

if args.decode:
    urldecode = True
else:
    urldecode = False

# print(fb_cookie)

def doMultiSearch( term, urldecode, page ):
    zero_result = 0
    for i in range(page-5,page-1):
        if i != page and i in page_history and page_history[i] == 0:
            zero_result = zero_result + 1

    if zero_result < 3:
        # print(page)
        s_results = goop.search( term, fb_cookie, page, True )
        # print(s_results)
        page_history[page] = len(s_results)
        for i in s_results:
            if urldecode:
                print( urllib.parse.unquote(s_results[i]['url']) )
            else:
                print( s_results[i]['url'] )
    else:
        for i in range(page,end_page):
            page_history[i] = 0


page_history = {}
start_page = 0
end_page = 100

pool = Pool( 5 )
pool.map( partial(doMultiSearch,search,urldecode), range(start_page,end_page) )
pool.close()
pool.join()
