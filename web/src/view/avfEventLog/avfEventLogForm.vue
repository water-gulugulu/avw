<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="区块编号:">
                <el-input v-model="formData.blockNumber" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="合约地址:">
                <el-input v-model="formData.contract" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="发起地址:">
                <el-input v-model="formData.form" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="接收地址:">
                <el-input v-model="formData.to" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="日志索引:">
                <el-input v-model="formData.index" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="事务索引:">
                <el-input v-model="formData.txIndex" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="事件名称:">
                <el-input v-model="formData.name" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="操作金额:"><el-input v-model.number="formData.number" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="tokens:">
                <el-input v-model="formData.tokens" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="时间戳:"><el-input v-model.number="formData.createTime" clearable placeholder="请输入"></el-input>
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
    createAvfEventLog,
    updateAvfEventLog,
    findAvfEventLog
} from "@/api/avfEventLog";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "AvfEventLog",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            blockNumber:"",
            contract:"",
            form:"",
            to:"",
            index:"",
            txIndex:"",
            name:"",
            number:0,
            tokens:"",
            createTime:0,
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createAvfEventLog(this.formData);
          break;
        case "update":
          res = await updateAvfEventLog(this.formData);
          break;
        default:
          res = await createAvfEventLog(this.formData);
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
    const res = await findAvfEventLog({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.reavfEventLog
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