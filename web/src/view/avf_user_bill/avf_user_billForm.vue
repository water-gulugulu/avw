<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="用户ID:"><el-input v-model.number="formData.uid" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="卡牌ID:"><el-input v-model.number="formData.cardId" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="钱包地址:">
                <el-input v-model="formData.address" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="账单类型:">
                <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.type" clearable ></el-switch>
          </el-form-item>
           
             <el-form-item label="金额:"><el-input v-model.number="formData.money" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="手续费:"><el-input v-model.number="formData.fees" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="余额:"><el-input v-model.number="formData.balance" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="收入支出:">
                <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.payment" clearable ></el-switch>
          </el-form-item>
           
             <el-form-item label="支付方式:">
                <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.payType" clearable ></el-switch>
          </el-form-item>
           
             <el-form-item label="描述:">
                <el-input v-model="formData.detail" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="创建时间:"><el-input v-model.number="formData.createTime" clearable placeholder="请输入"></el-input>
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
    createAvfUserBill,
    updateAvfUserBill,
    findAvfUserBill
} from "@/api/avf_user_bill";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "AvfUserBill",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            uid:0,
            cardId:0,
            address:"",
            type:false,
            money:0,
            fees:0,
            balance:0,
            payment:false,
            payType:false,
            detail:"",
            createTime:0,
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createAvfUserBill(this.formData);
          break;
        case "update":
          res = await updateAvfUserBill(this.formData);
          break;
        default:
          res = await createAvfUserBill(this.formData);
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
    const res = await findAvfUserBill({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.reavfUserBill
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