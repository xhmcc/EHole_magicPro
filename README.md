# EHole_magicPro

在EHole_magic原本的基础上新增获取quake资产，更新hunter资产获取

（不影响原版功能的使用）

### fofa识别

注意：从FOFA识别需要配置FOFA 密钥以及邮箱，在config.ini内配置好密钥以及邮箱即可使用。

```
ehole finger -s domain="baidu.com"  // 支持所有fofa语法
```

### hunter识别

注意：从hunter识别需要配置hunter 密钥，在config.ini内配置好密钥即可使用。

```
ehole finger -b ip="182.61.201.211"  // 支持所有hunter语法
```

### quake识别

注意：从quake识别需要配置quake 密钥，在config.ini内配置好quake API KEY即可使用。

```
ehole finger -k ip:"182.61.201.211"  // 支持所有quake语法
```

### 本地识别

```
ehole finger -l url.txt  // 从文件中加载url扫描
```

### 单个目标识别

```
ehole finger -u http://www.baidu.com // 单个url检测
```

## 感谢

EHole_magic: https://github.com/lemonlove7/EHole_magic

EHole：https://github.com/EdgeSecurityTeam/EHole

EHole-modify：https://github.com/A10nggg/EHole-modify/

## 免责声明

本工具仅作为安全研究交流，请勿用于非法用途。如您在使用本工具的过程中存在任何非法行为，您需自行承担相应后果，本人将不承担任何法律及连带责任。
