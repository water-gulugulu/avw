<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="用户的上级地址:">
                <el-input v-model="formData.pid" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="用户名:">
                <el-input v-model="formData.username" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="帐号手机号:">
                <el-input v-model="formData.mobile" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="钱包地址:">
                <el-input v-model="formData.walletAddress" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="密码:">
                <el-input v-model="formData.password" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="支付密码:">
                <el-input v-model="formData.payPassword" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="登录时间:"><el-input v-model.number="formData.loginTime" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="登录ip:">
                <el-input v-model="formData.loginIp" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="登录次数:"><el-input v-model.number="formData.loginTimes" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="创建时间:"><el-input v-model.number="formData.createdTime" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="状态:">
                <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.status" clearable ></el-switch>
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
    createAvfUser,
    updateAvfUser,
    findAvfUser
} from "@/api/avf_user";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "AvfUser",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            pid:"",
            username:"",
            mobile:"",
            walletAddress:"",
            password:"",
            payPassword:"",
            loginTime:0,
            loginIp:"",
            loginTimes:0,
            createdTime:0,
            status:false,
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createAvfUser(this.formData);
          break;
        case "update":
          res = await updateAvfUser(this.formData);
          break;
        default:
          res = await createAvfUser(this.formData);
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
    const res = await findAvfUser({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.reavfUser
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