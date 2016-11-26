let user = {
  get name () {
    return this.$name
  },
  set name (value) {
    this.$name = value
    console.log(`name change to ${value}`)
  }
}

user.name = 'pat'
