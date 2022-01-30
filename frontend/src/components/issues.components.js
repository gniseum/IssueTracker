import React, {useState, useEffect} from 'react';
import axios from 'axios';
import { Button, Form, Container, Modal } from 'react-bootstrap';
import Issue from './single-issue.component';

const Issues = () => {
    const [issues, setIssues] = useState([])
    const [refreshData, setRefreshData] = useState(false)

    const [changeIssue, setChangeIssue] = useState({"change": false, "id": 0})
    const [changeTaskedUser, setChangeTaskedUser] = useState({"change": false,  "id": 0})
    const [newTaskedUserName, setNewTaskedUserName] = useState("")

    const [addNewIssue, setAddNewIssue] = useState(false)
    const [newIssue, setNewIssue] = useState({"dish": "", "server":
        "", "table": 0, "price": 0})

    // Run a request for all issues
    useEffect(() => {
        getAllIssues();
    }, [])

    // Refresh the page
    if(refreshData) {
        setRefreshData(false);
        getAllIssues();
    }
    return (
 
    );

    // Change taskedUser
    function changeTaskedUserForIssue() {
        changeTaskedUser.change = false
        var url = "http://127.0.0.1:5000/taskedUser/update" + changeTaskedUser.id
        axios.put(url, {"server": newTaskedUserName}).then(
            response => {
                console.log(response.status)
                if(response.status === 200) {
                    setRefreshData(true)
                }
            }
        )
    }

    // Change issue
    function changeSingleIssue() {
        changeIssue.change = false
        var url = "http://127.0.0.1:5000/issue/update" + changeIssue.id
        axios.put(url, newIssue).then(
            response => {
                console.log(response.status)
                if(response.status === 200) {
                    setRefreshData(true)
                }
            }
        )
    }


    // Create new issue
    function addSingleIssue() {
        setAddNewIssue(false)
        var url = "http://127.0.0.1:5000/issue/create"
        axios.post(url, {
            "server": newIssue.server,
            "dish": newIssue.dish,
            "table": newIssue.table,
            "price": parseFloat(newIssue.price)
        }).then(
            response => {
                console.log(response.status)
                if(response.status === 200) {
                    setRefreshData(true)
                }
            }
        )
    }

    // Get all issues
    function getAllIssues() {
        var url = "http://127.0.0.1:5000/issues"
        axios.get(url, {
            responseType: 'json'
        }).then(
            response => {
                console.log(response.status)
                if(response.status === 200) {
                    setRefreshData(true)
                }
            }
        )
    }

    // Delete an issue
    function deleteSingleIssue() {
        var url = "http://127.0.0.1:5000/issue/delete" + changeIssue.id
        axios.delete(url, {}).then(
            response => {
                console.log(response.status)
                if(response.status === 200) {
                    setRefreshData(true)
                }
            }
        )
    }
 }

export default Issues
