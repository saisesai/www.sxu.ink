import axios from "axios";
import {ElMessage} from "element-plus/es";

export type ArtworkInfo = {
    ti_ming: string, cover: string, zang_nian: string, chao_dai: string
    chu_tu_di: string, xian_cang_di: string, hang_zi_shu: string,
    chi_cun: string, shuo_ming: string, zu_nian: string,
    nian_ling: string, xing_bie: string, ji_guan: string,
    zhi_gai: string, ming_wen: string,
    pages: string
}

export function NewArtwork(value: ArtworkInfo): Promise<string> {
    return new Promise<string>((resolve) => {
        axios.post("/api/artwork/new").then(resp => {
            let uuid = resp.data["uuid"] as string;
            axios.post("/api/artwork/set/" + uuid, value).then(() => {
                ElMessage.success("作品已保存！");
                resolve(uuid);
            }).catch(() => {
                ElMessage.error("保存作品失败！");
            })
        }).catch(() => {
            ElMessage.error("创建作品失败！");
        })
    })
}

export function SoftDeleteArtwork(uuid: string) {
    let now = new Date();
    axios.post("/api/artwork/set/" + uuid,
        {"delete_at": now.toUTCString()}).then(() => {
        ElMessage.success("作品已删除！");
    }).catch((err) => {
        ElMessage.error("删除作品失败！");
        console.log(err)
    })
}

export function GetArtworkInfo(uuid: string, value: ArtworkInfo): Promise<void> {
    return new Promise<void>(resolve => {
        axios.post("/api/artwork/get/" + uuid).then(resp => {
            value.chao_dai = resp.data["chao_dai"];
            value.ti_ming = resp.data["ti_ming"];
            value.cover = resp.data["cover"];
            value.zang_nian = resp.data["zang_nian"];
            value.chu_tu_di = resp.data["chu_tu_di"];
            value.xian_cang_di = resp.data["xian_cang_di"];
            value.hang_zi_shu = resp.data["hang_zi_shu"];
            value.chi_cun = resp.data["chi_cun"];
            value.shuo_ming = resp.data["shuo_ming"];
            value.zu_nian = resp.data["zu_nian"];
            value.nian_ling = resp.data["nian_ling"];
            value.xing_bie = resp.data["xing_bie"];
            value.ji_guan = resp.data["ji_guan"];
            value.zhi_gai = resp.data["zhi_gai"];
            value.ming_wen = resp.data["ming_wen"];
            value.pages = resp.data["pages"];
            resolve();
        }).catch(() => {
            ElMessage.error("获取作品数据失败！");
        })
    })
}

export function SetArtwork(uuid: string, value: ArtworkInfo) {
    axios.post("/api/artwork/set/" + uuid, value).then(() => {
        ElMessage.success("作品已保存！");
    }).catch((err) => {
        ElMessage.error("保存作品失败！");
        console.log(err)
    })
}

export function ParsePages(page_str: string): string[] {
    page_str = page_str.replace(/\s/g, '');
    return page_str.split(";").filter(n => n);
}