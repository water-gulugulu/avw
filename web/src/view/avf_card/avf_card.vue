<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">                        
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增卡牌</el-button>
        </el-form-item>
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
    <el-table-column label="日期" width="180">
         <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
    </el-table-column>
    
    <el-table-column label="卡牌名称" prop="name" width="120"></el-table-column> 
    
    <el-table-column label="星力值" prop="star" width="120"></el-table-column> 
    
    <el-table-column label="卡牌价格" prop="money" width="120"></el-table-column> 
    
    <el-table-column label="支付钱包地址" prop="walletAddress" width="120"></el-table-column> 
    
    <el-table-column label="合约地址" prop="contractAddress" width="120"></el-table-column> 
    
    <el-table-column label="发行数量" prop="number" width="120"></el-table-column> 
    
    <el-table-column label="等级" prop="level" width="120"></el-table-column> 
    
    <el-table-column label="作者" prop="author" width="120"></el-table-column> 
    
    <el-table-column label="描述" prop="desc" width="120"></el-table-column> 
    
    <el-table-column label="状态" prop="status" width="120">
         <template slot-scope="scope">{{scope.row.status|formatBoolean}}</template>
    </el-table-column>
    
    <el-table-column label="创建时间" prop="createDate" width="120"></el-table-column> 
    
    <el-table-column label="修改时间" prop="updateDate" width="120"></el-table-column> 
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateAvfCard(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
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


          <el-form-item label="等级" prop="level">
              <el-cascader
                      :disabled="dialogType=='add'"
                      :options="AuthorityOption"
                      :props="{ checkStrictly: true,label:'authorityName',value:'authorityId',disabled:'disabled',emitPath:false}"
                      :show-all-levels="false"
                      filterable
                      v-model="formData.level"
              ></el-cascader>
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
    createAvfCard,
    deleteAvfCard,
    deleteAvfCardByIds,
    updateAvfCard,
    findAvfCard,
    getAvfCardList
} from "@/api/avf_card";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "AvfCard",
  mixins: [infoList],
  data() {
    return {
        AuthorityOption: [
            {
                authorityId: 1,
                authorityName: "N"
            },
            {
                authorityId: 2,
                authorityName: "R"
            },
            {
                authorityId: 3,
                authorityName: "SR"
            },
            {
                authorityId: 4,
                authorityName: "SSR"
            },
        ],
      listApi: getAvfCardList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
            name:"",
            star:0,
            money:0,
            walletAddress:"",
            contractAddress:"",
            number:0,
            level:0,
            author:"",
            desc:"",
            status:false,
            createDate:new Date(),
            updateDate:new Date(),
            
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
           this.deleteAvfCard(row);
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
        const res = await deleteAvfCardByIds({ ids })
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
    async updateAvfCard(row) {
      const res = await findAvfCard({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reavfCard;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          name:"",
          star:0,
          money:0,
          walletAddress:"",
          contractAddress:"",
          number:0,
          level:0,
          author:"",
          desc:"",
          status:false,
          createDate:new Date(),
          updateDate:new Date(),
          
      };
    },
    async deleteAvfCard(row) {
      const res = await deleteAvfCard({ ID: row.ID });
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
