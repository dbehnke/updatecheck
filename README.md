# updatecheck

Updatecheck is a project to scrape websites for current version and download url of a piece of software.

It is certainly not limited to that, you can add more fields as you'd like.  See examples in the definitions folder. 

## quickstart

```bash
$ go build -mod vendor
$ ./updatecheck results definitions results.json
$ cat results.json
```

## contributions

1. you need go 1.11+ (modules)

2. build updatecheck, see quickstart above.

3. pick existing .yml or create a new one.

4. test it out 

```bash
$ ./updatecheck results definitions/nameofyour.yml results.json
$ cat results.json
```

5. submit a pull request

## sample results.json

```json
{
  "adobeflashplayer": {
    "mac": {
      "current_version": "31.0.0.108",
      "download_url": "https://fpdownload.macromedia.com/get/flashplayer/pdc/31.0.0.108/install_flash_player_osx.dmg"
    },
    "windows": {
      "current_version": "31.0.0.108",
      "download_url": "https://fpdownload.macromedia.com/get/flashplayer/pdc/31.0.0.108/install_flash_player.exe"
    }
  },
  "chrome-stable": {
    "mac": {
      "current_version": "69",
      "download_url": "https://dl.google.com/chrome/mac/stable/GGRO/googlechrome.dmg"
    }
  },
  "firefox": {
    "mac": {
      "current_version": "62.0.2",
      "download_url": "https://download-installer.cdn.mozilla.net/pub/firefox/releases/62.0.2/mac/en-US/Firefox%2062.0.2.dmg"
    },
    "windows": {
      "current_version": "62.0.2",
      "download_url": "https://download-installer.cdn.mozilla.net/pub/firefox/releases/62.0.2/win32/en-US/Firefox%20Setup%2062.0.2.exe"
    }
  },
  "firefoxesr": {
    "mac": {
      "current_version": "60.2.1",
      "download_url": "https://download-installer.cdn.mozilla.net/pub/firefox/releases/60.2.1esr/mac/en-US/Firefox%2060.2.1esr.dmg"
    },
    "windows": {
      "current_version": "60.2.1",
      "download_url": "https://download-installer.cdn.mozilla.net/pub/firefox/releases/60.2.1esr/win32/en-US/Firefox%20Setup%2060.2.1esr.exe"
    }
  },
  "gimp": {
    "mac": {
      "current_version": "2.10.6",
      "download_url": "https://download.gimp.org/mirror/pub/gimp/v2.10/osx/gimp-2.10.6-x86_64.dmg"
    },
    "windows": {
      "current_version": "2.10.6",
      "download_url": "https://download.gimp.org/mirror/pub/gimp/v2.10/windows/gimp-2.10.6-setup.exe"
    }
  },
  "keepassxc": {
    "mac": {
      "current_version": "2.3.4",
      "download_url": "https://github.com/keepassxreboot/keepassxc/releases/download/2.3.4/KeePassXC-2.3.4.dmg"
    }
  },
  "libreoffice": {
    "mac": {
      "current_version": "6.1.2",
      "download_url": "http://download.documentfoundation.org/libreoffice/stable/6.1.2/mac/x86_64/LibreOffice_6.1.2_MacOS_x86-64.dmg"
    },
    "windows": {
      "current_version": "6.1.2",
      "download_url": "http://download.documentfoundation.org/libreoffice/stable/6.1.2/win/x86_64/LibreOffice_6.1.2_Win_x64.msi"
    }
  },
  "libreoffice-enterprise": {
    "mac": {
      "current_version": "6.0.6",
      "download_url": "http://download.documentfoundation.org/libreoffice/stable/6.0.6/mac/x86_64/LibreOffice_6.0.6_MacOS_x86-64.dmg"
    },
    "windows": {
      "current_version": "6.0.6",
      "download_url": "http://download.documentfoundation.org/libreoffice/stable/6.0.6/win/x86_64/LibreOffice_6.0.6_Win_x64.msi"
    }
  },
  "oraclejavajdk": {
    "mac": {
      "current_version": "8u181-b13",
      "download_url": "http://download.oracle.com/otn-pub/java/jdk/8u181-b13/96a7b8442fe848ef90c96a2fad6ed6d1/jdk-8u181-macosx-x64.dmg"
    },
    "windows": {
      "current_version": "8u181-b13",
      "download_url": "http://download.oracle.com/otn-pub/java/jdk/8u181-b13/96a7b8442fe848ef90c96a2fad6ed6d1/jdk-8u181-windows-x64.exe"
    }
  },
  "thunderbird": {
    "mac": {
      "current_version": "60.0",
      "download_url": "https://download-installer.cdn.mozilla.net/pub/thunderbird/releases/60.0/mac/en-US/Thunderbird%2060.0.dmg"
    },
    "windows": {
      "current_version": "60.0",
      "download_url": "https://download-installer.cdn.mozilla.net/pub/thunderbird/releases/60.0/win32/en-US/Thunderbird%20Setup%2060.0.exe"
    }
  },
  "tigervnc": {
    "mac": {
      "current_version": "1.9.0",
      "download_url": "https://bintray.com/tigervnc/stable/download_file?file_path=TigerVNC-1.9.0.dmg"
    },
    "windows": {
      "current_version": "1.9.0",
      "download_url": "https://bintray.com/tigervnc/stable/download_file?file_path=vncviewer-1.9.0.exe"
    }
  },
  "virtualbox": {
    "mac": {
      "current_version": "5.2.18",
      "download_url": "https://download.virtualbox.org/virtualbox/5.2.18/VirtualBox-5.2.18-124319-OSX.dmg"
    },
    "windows": {
      "current_version": "5.2.18",
      "download_url": "https://download.virtualbox.org/virtualbox/5.2.18/VirtualBox-5.2.18-124319-Win.exe"
    }
  },
  "vscode": {
    "linux-rpm": {
      "current_version": "1.27.2",
      "download_url": "https://vscode-update.azuresites.net/v1.27.2/linux-rpm-x64/stable"
    },
    "mac": {
      "current_version": "1.27.2",
      "download_url": "https://vscode-update.azurewebsites.net/1.27.2/darwin/stable"
    },
    "windows": {
      "current_version": "1.27.2",
      "download_url": "https://vscode-update.azuresites.net/v1.27.2/win32-x64-user/stable"
    }
  }
}
```

