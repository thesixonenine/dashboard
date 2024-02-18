import request from "@/http/request";
import type {Hospital} from "@/api/home_type.ts";
import {Content, HospitalDict} from "@/api/home_type.ts";
import {Response} from "@/api/common_type.ts";

enum Api {
    card_list = "/hosp/hospital/{pageNo}/{pageSize}",
    card_level = '/cmn/dict/findByDictCode/{dictCode}',
    card_by_name = '/hosp/hospital/findByHosname/{hosname}',
}

export const cardListReq = (page: number, pageSize: number, hostype: string = '', districtCode: string = '') => {
    return request.get<any, Response<Hospital>>(Api.card_list.replace('{pageNo}', String(page))
            .replace('{pageSize}', String(pageSize)) +
        "?hostype=" + hostype + "&districtCode=" + districtCode);
};

export const cardLevelReq = (dictCode: string) => {
    return request.get<any, Response<HospitalDict[]>>(Api.card_level.replace('{dictCode}', dictCode));
};

export const cardByNameReq = (name: string) => {
    return request.get<any, Response<Content[]>>(Api.card_by_name.replace('{hosname}', name));
};
