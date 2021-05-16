<template>
  <div>
    <el-row :gutter="15">
      <el-form ref="blind_box" :model="formData" :rules="rules" size="mini" label-width="100px"
        label-position="top">
        <el-col :span="9">
          <el-form-item label="一连抽" prop="one">
            <el-input v-model="formData.one" placeholder="请输入盲盒一连抽价格" :maxlength="9" show-word-limit clearable
              prefix-icon='el-icon-money' :style="{width: '100%'}">
              <template slot="prepend">价格</template>
              <template slot="append">HT</template>
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="10">
          <el-form-item label="十连抽" prop="ten">
            <el-input v-model="formData.ten" placeholder="请输入盲盒十连抽价格" :maxlength="9" show-word-limit clearable
              prefix-icon='el-icon-money' :style="{width: '100%'}">
              <template slot="prepend">价格</template>
              <template slot="append">HT</template>
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item size="large">
            <el-button type="primary" @click="submitForm">提交</el-button>
            <el-button @click="resetForm">重置</el-button>
          </el-form-item>
        </el-col>
      </el-form>
    </el-row>
  </div>
</template>
<script>
export default {
  components: {},
  props: [],
  data() {
    return {
      formData: {
        one: 100,
        ten: 950,
      },
      rules: {
        one: [{
          required: true,
          message: '请输入盲盒一连抽价格',
          trigger: 'blur'
        }],
        ten: [{
          required: true,
          message: '请输入盲盒十连抽价格',
          trigger: 'blur'
        }],
      },
    }
  },
  computed: {},
  watch: {},
  created() {},
  mounted() {},
  methods: {
    submitForm() {
      this.$refs['blind_box'].validate(valid => {
        if (!valid) return
        // TODO 提交表单
      })
    },
    resetForm() {
      this.$refs['blind_box'].resetFields()
    },
  }
}

</script>
<style>
</style>
