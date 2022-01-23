# vk-archive-assets-downloader

Utility for downloading attached files from the [VKontakte data archive](https://vk.com/data_protection?section=rules).

## Quick start:

1. Request and download archive on [this page](https://vk.com/data_protection?section=rules)
2. Unarchive to a folder
3. Run `./vk-archive-assets-downloader --src='~/downloads/archive'`

## Parameters
``` bash
./vk-archive-assets-downloader --src=~/downloads/archive --dialogs=259460211,359460212 --types=Photo,Video -dest ~/dest
# OR
./vk-archive-assets-downloader -src ~/downloads/archive -dialogs 259460211,359460212 -types Photo,Video -dest ~/dest
```

| Parameter | Default value | Description |
|---|---|---|
| --src | `./archive` | Path to archive folder |
| --dialogs | `N/A` | Comma separated string with dialog ids. Dialog id you can found in dialog URL (sel query parameter) `https://vk.com/im?sel=<Dialog Id>` |
| -types | `N/A` |  Comma separated string with the attachment type name. A type must be specified in the language of the archive. Example `-types Photo,Video` for english and `-types Фотография,Видеозапись` for russian |
| -dest | `./dest` | Path to destination folder. |
| -encoding | `Windows1251` | Archive source file encoding. [The following encodings](#supported encodings) are supported. |
| -pool | 1000 | Length of attachments queue |
| -threads | *Number of logical CPUs usable by the current process* | Number of threads to download attachments |

## Supported encodings
- `CodePage037` is this IBM Code Page 037 encoding.
- `CodePage1047` is this IBM Code Page 1047 encoding.
- `CodePage1140` is this IBM Code Page 1140 encoding.
- `CodePage437` is this IBM Code Page 437 encoding.
- `CodePage850` is this IBM Code Page 850 encoding.
- `CodePage852` is this IBM Code Page 852 encoding.
- `CodePage855` is this IBM Code Page 855 encoding.
- `CodePage858` is this Windows Code Page 858 encoding.
- `CodePage860` is this IBM Code Page 860 encoding.
- `CodePage862` is this IBM Code Page 862 encoding.
- `CodePage863` is this IBM Code Page 863 encoding.
- `CodePage865` is this IBM Code Page 865 encoding.
- `CodePage866` is this IBM Code Page 866 encoding.
- `ISO8859_1` is this ISO 8859-1 encoding.
- `ISO8859_10` is this ISO 8859-10 encoding.
- `ISO8859_13` is this ISO 8859-13 encoding.
- `ISO8859_14` is this ISO 8859-14 encoding.
- `ISO8859_15` is this ISO 8859-15 encoding.
- `ISO8859_16` is this ISO 8859-16 encoding.
- `ISO8859_2` is this ISO 8859-2 encoding.
- `ISO8859_3` is this ISO 8859-3 encoding.
- `ISO8859_4` is this ISO 8859-4 encoding.
- `ISO8859_5` is this ISO 8859-5 encoding.
- `ISO8859_6` is this ISO 8859-6 encoding.
- `ISO8859_7` is this ISO 8859-7 encoding.
- `ISO8859_8` is this ISO 8859-8 encoding.
- `ISO8859_9` is this ISO 8859-9 encoding.
- `KOI8R` is this KOI8-R encoding.
- `KOI8U` is this KOI8-U encoding.
- `Macintosh` is this Macintosh encoding.
- `MacintoshCyrillic` is this Macintosh Cyrillic encoding.
- `Windows1250` is this Windows 1250 encoding.
- `Windows1251` is this Windows 1251 encoding.
- `Windows1252` is this Windows 1252 encoding.
- `Windows1253` is this Windows 1253 encoding.
- `Windows1254` is this Windows 1254 encoding.
- `Windows1255` is this Windows 1255 encoding.
- `Windows1256` is this Windows 1256 encoding.
- `Windows1257` is this Windows 1257 encoding.
- `Windows1258` is this Windows 1258 encoding.
- `Windows874` is this Windows 874 encoding.
