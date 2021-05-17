<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="卡牌出售价格:"><el-input v-model.number="formData.price" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="手续费:"><el-input v-model.number="formData.fees" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="状态:"><el-input v-model.number="formData.status" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="卡牌名称:">
                <el-input v-model="formData.cardName" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="交易hash:">
                <el-input v-model="formData.txHash" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="出售人:">
                <el-input v-model="formData.from" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="购买人:">
                <el-input v-model="formData.to" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="区块号:">
                <el-input v-model="formData.block" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="手续费地址:">
                <el-input v-model="formData.system" clearable placeholder="请输入" ></el-input>
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
    createAvfCardTransfer,
    updateAvfCardTransfer,
    findAvfCardTransfer
} from "@/api/avf_card_transfer";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "AvfCardTransfer",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            price:0,
            fees:0,
            status:0,
            cardName:"",
            txHash:"",
            from:"",
            to:"",
            block:"",
            system:"",
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createAvfCardTransfer(this.formData);
          break;
        case "update":
          res = await updateAvfCardTransfer(this.formData);
          break;
        default:
          res = await createAvfCardTransfer(this.formData);
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
    const res = await findAvfCardTransfer({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.reavfCardTransfer
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