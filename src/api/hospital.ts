import request from "@/http/request";
import {Response} from "@/api/common_type.ts";
import {HospitalDetail} from "@/api/hospital_type.ts";

enum Api {
    hospital_detail = '/hosp/hospital/{hoscode}',
}

export const cardDetailReq = (hoscode: string) => {
    return request.get<any, Response<HospitalDetail>>(Api.hospital_detail.replace('{hoscode}', hoscode))
}