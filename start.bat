@echo off
echo	echo Usage:	
echo	ehole finger [flags]
echo.
echo [*] options:
echo.
echo     -u        识别单个目标
echo     -f        从fofa提取资产，进行指纹识别，仅仅支持ip或者ip段
echo     -s        从fofa提取资产，进行指纹识别，支持fofa所有语法
echo     -a        从hunter提取资产，进行指纹识别，仅仅支持ip或者ip段
echo     -b        从hunter提取资产，进行指纹识别，支持hunter所有语法
echo     -q        从quake提取资产，进行指纹识别，仅仅支持ip或者ip段
echo     -k        从quake提取资产，进行指纹识别，支持quake所有语法
echo     -p        指定访问目标时的代理，支持http代理和socks5
echo     -t        指纹识别线程大小。 (default 100)
echo.
echo [*] useage:
echo     ehole.exe finger -u http://baidu.com
echo.
call cmd.exe