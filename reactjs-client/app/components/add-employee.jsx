import React, { Component, PropTypes } from 'react'

export default class AddEmployee extends React.Component {

   render() {
      return (
         <div>
            <p>First Name: <input type='fname' ref='fname' /></p>
            <p>Last Name: <input type='lname' ref='lname' /></p>
            <button onClick={(e) => this.handleClick(e)}>
               Add
            </button>
         </div>
      )

   }
   handleClick(e) {
      const fname = this.refs.fname.value.trim()
      const lname = this.refs.lname.value.trim()
      this.props.addEmployee(fname, lname)
      fname.value = ''
      lname.value = ''
   }
}
