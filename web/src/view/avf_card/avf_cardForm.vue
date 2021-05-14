<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="卡牌名称:">
                <el-input v-model="formData.name" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="星力值:"><el-input v-model.number="formData.star" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="卡牌价格:">
                  <el-input-number v-model="formData.money" :precision="2" clearable></el-input-number>
           </el-form-item>
           
             <el-form-item label="支付钱包地址:">
                <el-input v-model="formData.walletAddress" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="合约地址:">
                <el-input v-model="formData.contractAddress" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="发行数量:"><el-input v-model.number="formData.number" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="等级:"><el-input  clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="作者:">
                <el-input v-model="formData.author" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="描述:">
                <el-input v-model="formData.desc" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="状态:">
                <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.status" clearable ></el-switch>
          </el-form-item>
           
             <el-form-item label="创建时间:">
                  <el-date-picker type="date" placeholder="选择日期" v-model="formData.createDate" clearable></el-date-picker>
           </el-form-item>
           
             <el-form-item label="修改时间:">
                  <el-date-picker type="date" placeholder="选择日期" v-model="formData.updateDate" clearable></el-date-picker>
           </el-form-item>
           <el-form-item>
           <el-button @click="save" type="primary">保存</el-button>
           <el-button @click="back" type="primary">返回</el-button>
           </el-form-item>
    </el-form>
</div>
</template>

<script>
import {
    createAvfCard,
    updateAvfCard,
    findAvfCard
} from "@/api/avf_card";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "AvfCard",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            name:"",
            star:0,
            money:0,
            walletAddress:"",
            contractAddress:"",
            number:0,
            level:'',
            author:"",
            desc:"",
            status:false,
            createDate:new Date(),
            updateDate:new Date(),
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createAvfCard(this.formData);
          break;
        case "update":
          res = await updateAvfCard(this.formData);
          break;
        default:
          res = await createAvfCard(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
      }
    },
    back(){
        this.$router.go(-1)
    }
  },
  async created() {
   // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if(this.$route.query.id){
    const res = await findAvfCard({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.reavfCard
       this.type = "update"
     }
    }else{
     this.type = "create"
   }
  
}
};
</script>

<style>
</style>