// 医院详情

export interface Hospital {
    content: Content[],
    pageable: Pageable,
    totalPages: number,
    totalElements: number,
    last: boolean,
    first: boolean,
    sort: Sort,
    numberOfElements: number,
    size: number,
    number: number,
    empty: boolean,
}

export interface Param {
    hostypeString: string,
    fullAddress: string,
}

export interface BookingRule {
    cycle: number,
    releaseTime: string,
    stopTime: string,
    quitDay: number,
    quitTime: string,
    rule: string[],
}

export interface Content {
    id: string,
    createTime: string,
    updateTime: string,
    isDeleted: number,
    param: Param,
    hoscode: string,
    hosname: string,
    hostype: string,
    provinceCode: string,
    cityCode: string,
    districtCode: string,
    address: string,
    logoData: string,
    intro: string,
    route: string,
    status: number,
    bookingRule: BookingRule,
}

export interface Sort {
    sorted: boolean,
    unsorted: boolean,
    empty: boolean,
}

export interface Pageable {
    sort: Sort,
    pageNumber: number,
    pageSize: number,
    offset: number,
    paged: boolean,
    unpaged: boolean,
}

export interface Sort {
    sorted: boolean,
    unsorted: boolean,
    empty: boolean,
}


// 医院等级和区域字典

export interface HospitalDict {
    id: number,
    createTime: string,
    updateTime: string,
    isDeleted: number,
    parentId: number,
    name: string,
    value: string,
    dictCode: string,
    hasChildren: boolean,
}
