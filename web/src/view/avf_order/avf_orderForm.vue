<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="用户ID:"><el-input v-model.number="formData.uid" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="订单编号:">
                <el-input v-model="formData.orderSn" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="支付价格:">
                  <el-input-number v-model="formData.price" :precision="2" clearable></el-input-number>
           </el-form-item>
           
             <el-form-item label="购买数量:"><el-input v-model.number="formData.num" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="剩余数量:"><el-input v-model.number="formData.number" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="订单状态:">
                <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.status" clearable ></el-switch>
          </el-form-item>
           
             <el-form-item label="支付时间:"><el-input v-model.number="formData.payTime" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="事务哈希:">
                <el-input v-model="formData.txHash" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="区块编号:">
                <el-input v-model="formData.block" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="手续费:">
                <el-input v-model="formData.gas" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="手续费价格:">
                <el-input v-model="formData.gasPrice" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="支付地址:">
                <el-input v-model="formData.from" clearable placeholder="请输入" ></el-input>
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
    createAvfOrder,
    updateAvfOrder,
    findAvfOrder
} from "@/api/avf_order";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "AvfOrder",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            uid:0,
            orderSn:"",
            price:0,
            num:0,
            number:0,
            status:false,
            payTime:0,
            txHash:"",
            block:"",
            gas:"",
            gasPrice:"",
            from:"",
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createAvfOrder(this.formData);
          break;
        case "update":
          res = await updateAvfOrder(this.formData);
          break;
        default:
          res = await createAvfOrder(this.formData);
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
    const res = await findAvfOrder({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.reavfOrder
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