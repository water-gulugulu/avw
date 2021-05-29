<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="用户的上级地址">
          <el-input placeholder="搜索条件" v-model="searchInfo.pid"></el-input>
        </el-form-item>    
        <el-form-item label="用户名">
          <el-input placeholder="搜索条件" v-model="searchInfo.username"></el-input>
        </el-form-item>      
        <el-form-item label="钱包地址">
          <el-input placeholder="搜索条件" v-model="searchInfo.walletAddress"></el-input>
        </el-form-item>                  
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <!--<el-form-item>
          <el-button @click="openDialog" type="primary">新增用户</el-button>
        </el-form-item> -->
        <el-form-item>
          <el-popover placement="top" v-model="deleteVisible" width="160">
            <p>确定要删除吗？</p>
              <div style="text-align: right; margin: 0">
                <el-button @click="deleteVisible = false" size="mini" type="text">取消</el-button>
                <el-button @click="onDelete" size="mini" type="primary">确定</el-button>
              </div>
            <el-button icon="el-icon-delete" size="mini" slot="reference" type="danger">批量删除</el-button>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      :data="tableData"
      @selection-change="handleSelectionChange"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
    <el-table-column type="selection" width="55"></el-table-column>
    <el-table-column label="用户的上级地址" prop="pid" width="350">
        <template slot-scope="scope">{{ scope.row.pid}}</template>
    </el-table-column>
    
    <el-table-column label="用户名" prop="username" width="160"></el-table-column>
    
    <!--<el-table-column label="帐号手机号" prop="mobile" width="120"></el-table-column> -->
    
    <el-table-column label="钱包地址" prop="walletAddress" width="350"></el-table-column>
    
    <!-- <el-table-column label="密码" prop="password" width="120"></el-table-column> -->
    
    <!-- <el-table-column label="支付密码" prop="payPassword" width="120"></el-table-column> -->

        <!-- <el-table-column label="登录时间" prop="loginTime" width="120"></el-table-column> -->
    
    <el-table-column label="登录ip" prop="loginIp" width="150"></el-table-column>
    
    <el-table-column label="登录次数" prop="loginTimes" width="120"></el-table-column>

        <el-table-column label="创建时间" width="180">
            <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
        </el-table-column>

        <!--<el-table-column label="创建时间" prop="CreatedAt" width="120"></el-table-column> -->

    <el-table-column label="状态" prop="status" width="120">
         <template slot-scope="scope">{{scope.row.status|formatBoolean}}</template>
    </el-table-column>
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <!--<el-button class="table-button" @click="updateAvfUser(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button> -->
          <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
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
       
       <!--  <el-form-item label="密码:">
            <el-input v-model="formData.password" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="支付密码:">
            <el-input v-model="formData.payPassword" clearable placeholder="请输入" ></el-input>
      </el-form-item> -->
       
       <!--  <el-form-item label="登录时间:"><el-input v-model.number="formData.loginTime" clearable placeholder="请输入"></el-input>
      </el-form-item>

         <el-form-item label="登录ip:">
            <el-input v-model="formData.loginIp" clearable placeholder="请输入" ></el-input>
      </el-form-item>

         <el-form-item label="登录次数:"><el-input v-model.number="formData.loginTimes" clearable placeholder="请输入"></el-input>
      </el-form-item>

         <el-form-item label="创建时间:"><el-input v-model.number="formData.createdTime" clearable placeholder="请输入"></el-input>
      </el-form-item> -->
       
         <el-form-item label="状态:">
            <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.status" clearable ></el-switch>
      </el-form-item>
       </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
    createAvfUser,
    deleteAvfUser,
    deleteAvfUserByIds,
    updateAvfUser,
    findAvfUser,
    getAvfUserList
} from "@/api/avf_user";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "AvfUser",
  mixins: [infoList],
  data() {
    return {
      listApi: getAvfUserList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
            pid:"",
            username:"",
            mobile:"",
            walletAddress:"",
            password:"",
            payPassword:"",
            loginTime:0,
            loginIp:"",
            loginTimes:0,
            CreatedAt:"0",
            status:false,
            
      }
    };
  },
  filters: {
    formatDate: function(time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    },
    formatBoolean: function(bool) {
      if (bool != null) {
        return bool ? "是" :"否";
      } else {
        return "";
      }
    }
  },
  methods: {
      //条件搜索前端看此方法
      onSubmit() {
        this.page = 1
        this.pageSize = 10                
        if (this.searchInfo.status==""){
          this.searchInfo.status=null
        }      
        this.getTableData()
      },
      handleSelectionChange(val) {
        this.multipleSelection = val
      },
      deleteRow(row){
        this.$confirm('确定要删除吗?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
           this.deleteAvfUser(row);
        });
      },
      async onDelete() {
        const ids = []
        if(this.multipleSelection.length == 0){
          this.$message({
            type: 'warning',
            message: '请选择要删除的数据'
          })
          return
        }
        this.multipleSelection &&
          this.multipleSelection.map(item => {
            ids.push(item.ID)
          })
        const res = await deleteAvfUserByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          if (this.tableData.length == ids.length) {
              this.page--;
          }
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async updateAvfUser(row) {
      const res = await findAvfUser({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reavfUser;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
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
          
      };
    },
    async deleteAvfUser(row) {
      const res = await deleteAvfUser({ ID: row.ID });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        if (this.tableData.length == 1) {
            this.page--;
        }
        this.getTableData();
      }
    },
    async enterDialog() {
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
        this.closeDialog();
        this.getTableData();
      }
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    }
  },
  async created() {
    await this.getTableData();
  
}
};
</script>

<style>
</style>
