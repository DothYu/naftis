// Task 定义了一些有关创建Task和通知Task执行状态的常量。
const Task = {
  command: {
    APPLY: 'apply',
    ROLLBACK: 'rollback'
  },
  status: {
    EXECUTING: 1,
    SUCCESS: 2,
    FAIL: 3,
    UNDEFINED: 4
  },
  varFormType: {
    STRING: 1,
    NUMBER: 2,
    PERCENTAGE: 3,
    SELECT: 4,
    DATETIME: 5
  },
  commandint: {
    APPLY: 1,
    ROLLBACK: 5
  }
}

export {
  Task
}
