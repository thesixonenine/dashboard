import {defineStore} from "pinia";
import {Response} from "@/api/common_type.ts";
import {HospitalDetail} from "@/api/hospital_type.ts";
import {cardDetailReq} from "@/api/hospital.ts";
import {HospitalDetailState} from "@/store/modules/interface";

const useDetailStore = defineStore('HospitalDetail', {
    state: (): HospitalDetailState => {
        return {
            hospitalDetail: ({} as HospitalDetail),
        }
    },
    actions: {
        async getHospitalDetail(hoscode: string) {
            const result: Response<HospitalDetail> = await cardDetailReq(hoscode);
            if (result.code == 200) {
                console.log('cardDetail:\n', result.data);
                this.hospitalDetail = result.data;
            }
        }
    },
    getters: {},
});
export default useDetailStore;
