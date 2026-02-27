from dataclasses import dataclass
from typing import Dict, List, Optional


@dataclass
class UserInfo:
    userId: int
    username: str
    email: str
    group: str
    isRealname: bool
    regTime: int
    status: int
    traffic: int
    usedProxies: int
    friendlyGroup: str
    maxProxies: int
    inBound: int
    outBound: int
    todaySigned: bool
    banReason: Optional[str] = None
    realnameTimes: int = 0
    vipExpireTime: Optional[int] = None


@dataclass
class SystemStatus:
    status: int
    remark: str


@dataclass
class Proxy:
    proxyId: int
    username: str
    proxyName: str
    proxyType: str
    isBanned: bool
    isDisabled: bool
    localIp: str
    localPort: int
    remotePort: int
    nodeId: int
    runId: str
    isOnline: bool
    domain: str
    lastStartTime: int
    lastCloseTime: int
    clientVersion: str
    proxyProtocolVersion: str
    useEncryption: bool
    useCompression: bool
    locations: str
    accessKey: str
    hostHeaderRewrite: str
    httpPlugin: str
    crtPath: str
    keyPath: str
    requestHeaders: str
    responseHeaders: str
    httpUser: str
    httpPassword: str
    transportProtocol: str


@dataclass
class NodeConnection:
    nodeId: int
    name: str
    hostname: str


@dataclass
class ProxyListResponse:
    proxies: List[Proxy]
    nodes: List[NodeConnection]


@dataclass
class Node:
    nodeId: int
    name: str
    hostname: str
    description: str
    token: str
    servicePort: int
    adminPort: int
    adminPass: str
    allowGroup: str
    allowPort: str
    allowType: str
    region: str
    bandwidth: str
    isOnline: bool
    isDisabled: bool
    totalTrafficIn: int
    totalTrafficOut: int
    upTime: int
    version: str
    donateId: int = 0
    donateUser: str = ""


@dataclass
class NodeWithLoad(Node):
    loadPercent: int = 0


@dataclass
class CreateProxyDataResponse:
    nodes: List[NodeWithLoad]
    groups: List["UserGroup"]
    currentGroup: str


@dataclass
class ProxyConfigResponse:
    config: str
    type: str


@dataclass
class NodeStatus:
    nodeId: int
    name: str
    totalTrafficIn: int
    totalTrafficOut: int
    onlineClient: int
    onlineProxy: int
    isOnline: bool
    version: str
    uptime: int
    curConns: int
    loadPercent: int


@dataclass
class Statistics:
    users: int
    nodes: int
    proxies: int
    traffic: int


@dataclass
class UserGroup:
    name: str
    friendlyName: str
    maxProxies: int
    baseTraffic: int
    outBound: int
    inBound: int


@dataclass
class StoreItem:
    type: str
    name: str
    price: float
    unit: str
    description: str
    enabled: bool
    discountEnabled: bool
    discountPrice: float
    discountStartTime: int
    discountEndTime: int
    currentPrice: float
    isDiscountActive: bool
    discountRemainingSeconds: int


@dataclass
class OperationLog:
    logId: int
    category: str
    details: str
    ipAddress: str
    status: str
    createdAt: str


@dataclass
class OperationLogList:
    data: List[OperationLog]
    total: int
    page: int
    pageSize: int
    totalPages: int


@dataclass
class OperationLogStats:
    monthCount: int
    todayCount: int
    totalCount: int
    weekCount: int


@dataclass
class Ads:
    adsId: int
    adsOwner: str
    adsUrl: str
    adsType: str
    adsContent: str
    adsImageUrl: str
    adsStartTime: int
    adsExpire: int
    renewalPrice: float
    adsPlacement: str
    adsClick: int
    adsImpression: int
    adsStatus: int
    adsReviewNote: str
    adsReviewer: str
    adsReviewTime: int
    adsSlotId: int
    adsCreatedTime: int


@dataclass
class AdCredit:
    creditId: int
    userId: int
    username: str
    slotId: int
    slotName: Optional[str] = None
    total: int = 0
    used: int = 0
    updateTime: int = 0
    expireTime: int = 0


@dataclass
class Order:
    orderId: str
    userId: int
    type: str
    amount: int
    months: int
    money: float
    status: int
    payType: str
    payURL: str
    payInfo: str
    payHTML: str
    payQRCode: str
    tradeNo: str
    couponCode: str
    adSlotType: str
    createTime: int
    updateTime: int


@dataclass
class NodeDonate:
    donateId: int
    username: str
    nodeName: str
    hostname: str
    description: str
    servicePort: int
    adminPort: int
    adminPass: str
    allowGroup: str
    allowPort: str
    allowType: str
    region: str
    bandwidth: str
    status: int
    rejectReason: str
    applyTime: int
    reviewTime: int
    nodeId: int


@dataclass
class EasyStartProxy:
    proxyId: int
    username: str
    proxyName: str
    proxyType: str
    isBanned: bool
    isDisabled: bool
    localIp: str
    localPort: int
    remotePort: int
    runId: str
    isOnline: bool
    domain: str
    lastStartTime: int
    lastCloseTime: int
    clientVersion: str
    proxyProtocolVersion: str
    useEncryption: bool
    useCompression: bool
    locations: str
    accessKey: str
    hostHeaderRewrite: str
    httpPlugin: str
    crtPath: str
    keyPath: str
    requestHeaders: Dict[str, str]
    httpUser: str
    httpPassword: str
    nodeAddr: str
    nodePort: int
    nodeToken: str


@dataclass
class NodeNameListItem:
    nodeId: int
    name: str
    hostname: str


@dataclass
class Product:
    productId: str
    system: str
    arch: str
    name: str
    desc: str
    path: str
    version: str
    isPublic: bool


@dataclass
class DownloadSource:
    id: int
    path: str
    name: str
