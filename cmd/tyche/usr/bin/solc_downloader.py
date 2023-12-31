import urllib.request
import os

# Solidity版本号列表和目标操作系统平台
VERSIONS = ["0.4.10", "0.4.9", "0.4.8", "0.4.7", "0.4.6", "0.4.5", "0.4.4", "0.4.3", "0.4.2", "0.4.1", "0.4.0", "0.3.6", "0.3.5", "0.3.4", "0.3.3", "0.3.2", "0.3.1", "0.3.0", "0.2.2", "0.2.1", "0.2.0", "0.1.7", "0.1.6", "0.1.5", "0.1.4", "0.1.3", "0.1.2", "0.1.1", "0.1.0"]
PLATFORM = "linux"

# 根据操作系统平台确定文件扩展名
if PLATFORM == "windows":
    EXTENSION = "exe"
elif PLATFORM in ["linux", "macos"]:
    EXTENSION = "zip"
else:
    print("Unsupported platform:", PLATFORM)
    exit(1)

# 循环下载所有版本的Solidity解析器文件，并将文件保存到对应的目录中
for version in VERSIONS:
    # 创建版本目录
    dirname = f"solidity_{version}"
    if not os.path.exists(dirname):
        os.mkdir(dirname)

    # 构建Solidity解析器下载链接
    url = f"https://github.com/ethereum/solidity/releases/download/v{version}/solc-static-linux"
    filename = f"solc-static-linux"

    # 下载文件
    print("Downloading", filename)
    urllib.request.urlretrieve(url, filename)

    # 将下载的文件移动到版本目录中
    os.rename(filename, os.path.join(dirname, filename))

print("Done!")