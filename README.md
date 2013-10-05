# XBMC NFO Tool

For a given local folder, interactively creates NFO files that simply link to themoviedb.org.

Can also highlight movie files without an NFO.

# Installation

    go build -o somwhere/in/your/path/xbmc-nfo-tool xbmc-nfo-tool.go

I'd say run the test suite too, but the tests are ninja.

# Usage

Simplest, in directory your want to check

  xbmc-nfo-tool

will produce a list of potential NFO files to create

Then to actually create them:

    $ xbmc-nfo-tool -mode=cli
    Need to Create NFO for: /tank/test/aliens.dvd.[tcw].nfo
    Unstripped: aliens.dvd.[tcw].nfo
    Stripped Title: `aliens`
    Found 20 Results
             0 ] =>  Aliens (http://themoviedb.org/movie/679)
             1 ] =>  Cowboys & Aliens (http://themoviedb.org/movie/49849)
             2 ] =>  Aliens vs. Avatars (http://themoviedb.org/movie/79582)
             3 ] =>  Superior Firepower: Making 'Aliens' (http://themoviedb.org/movie/70844)
             4 ] =>  Aliens - Sind sie unter uns (http://themoviedb.org/movie/60962)
             5 ] =>  Evil Aliens (http://themoviedb.org/movie/10177)
             6 ] =>  Illegal Aliens (http://themoviedb.org/movie/46329)
             7 ] =>  Mutant Aliens (http://themoviedb.org/movie/30863)
             8 ] =>  Monsters vs Aliens (http://themoviedb.org/movie/15512)
             9 ] =>  Aliens vs Predator: Requiem (http://themoviedb.org/movie/440)
             10 ] =>  Aliens on Crack (http://themoviedb.org/movie/28515)
             11 ] =>  Aliens in the Attic (http://themoviedb.org/movie/20856)
             12 ] =>  Aliens of the Deep (http://themoviedb.org/movie/22559)
             13 ] =>  Ancient Aliens Debunked (http://themoviedb.org/movie/139860)
             14 ] =>  Aliens Gone Wild (http://themoviedb.org/movie/58784)
             15 ] =>  Roswell: The Aliens Attack (http://themoviedb.org/movie/121988)
             16 ] =>  Aliens First Christmas (http://themoviedb.org/movie/81172)
             17 ] =>  The Aliens are Coming (http://themoviedb.org/movie/209443)
             18 ] =>  Ancient Aliens: Underwater Worlds (http://themoviedb.org/movie/156168)
             19 ] =>  Breakfast of Aliens (http://themoviedb.org/movie/124049)

    >>>Pick a Number:: [0]

It's pretty obvious we want the first one in this case, but sometimes you won't be sure.
Just click through the link and see if it's right. If there's only one choice, this
tool will pick it for you.

# Motivation

I have a bunch of movie that always get recognised incorrectly by XBMC.
I don't use MySQL for my XBMC library.
I like to update and change often - XBMCbuntu, OpenElec.
I use XBMC on multiple machines, a htpc, a rasberrypi, a couple of laptops.
I have the media on a read-only (at least read-only to the XBMC instances) SMB share.

I want to ensure that XBMC gets it right. To do that, you simply need to drop
"nfo" files with the url of a scraper (read *themoviedb.org*).

I started doing this manually. Then realised how stupid that was, so I wrote this.

# Ideas

I did think about making this web based and then we could pop film cover pictures to assist with the choosing.
This was enough for me though so I didn't implement it.

# License

I choose MIT which basically means "don't blame me if this software screws up". Sounds good to me.

[The MIT License (MIT)](http://opensource.org/licenses/MIT)

Copyright (c) 2013 Chris Walker

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.