<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="用户ID:"><el-input v-model.number="formData.uid" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="订单ID:"><el-input v-model.number="formData.orderId" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="卡牌ID:"><el-input v-model.number="formData.cardId" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="算力值:"><el-input v-model.number="formData.star" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="卡牌状态:"><el-input v-model.number="formData.status" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="创建时间:"><el-input v-model.number="formData.createTime" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="修改时间:"><el-input v-model.number="formData.updateTime" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="获得方式:"><el-input v-model.number="formData.giveType" clearable placeholder="请输入"></el-input>
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
    createAvfOrderCard,
    updateAvfOrderCard,
    findAvfOrderCard
} from "@/api/order_ard";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "AvfOrderCard",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            uid:0,
            orderId:0,
            cardId:0,
            star:0,
            status:0,
            createTime:0,
            updateTime:0,
            giveType:0,
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createAvfOrderCard(this.formData);
          break;
        case "update":
          res = await updateAvfOrderCard(this.formData);
          break;
        default:
          res = await createAvfOrderCard(this.formData);
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
    const res = await findAvfOrderCard({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.reavfOrderCard
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