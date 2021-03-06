import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import moment from 'moment';
import MoreInfo from './Table_info'
import { DefaultApi } from '../../api/apis';
import { EntPatientrights } from '../../api/models/EntPatientrights';



const useStyles = makeStyles(theme => ({
  table: {
    minWidth: 650,
  },
  formControl: {
    margin: theme.spacing(3),
    width: 200,
  },
  root: {
    // width: '100%',
    maxWidth: 360,
    minWidth: 360,

    backgroundColor: theme.palette.background.paper,

  },
  div: {
    flexGrow: 1,
    display: 'flex',
    justifyContent: 'center',
  },
  section1: {
    margin: theme.spacing(3, 2),
  },
  section3: {
    margin: theme.spacing(3, 1, 1),
  },
}));



export default function ComponentsTable(sim: any) {
  const classes = useStyles();
  const api = new DefaultApi();









  var p = 0;
  console.log("sim", sim.sim)

  if (sim.sim.length > 0) {
    p = p + 1
  }





  if (p !== 0) {
    return (

      <TableContainer component={Paper}>
        <Table className={classes.table} aria-label="simple table">
          <TableHead>
            <TableRow>
              <TableCell align="center">No.</TableCell>
              <TableCell align="center">ชื่อ ประกัน</TableCell>
              <TableCell align="center">Medicalrecordstaff</TableCell>
              <TableCell align="center">ความสามารถสิทธิ์</TableCell>
              <TableCell align="center">ผู้ป่วย</TableCell>
              <TableCell align="center">ผู้ดูแล</TableCell>
              <TableCell align="center">วันที่สร้าง</TableCell>
              <TableCell align="center">เพิ่มเติม</TableCell>

            </TableRow>
          </TableHead>
          <TableBody>
            {sim.sim === undefined
              ? null
              : sim.sim.map((item: any) => (
                <TableRow>
                  <TableCell align="center">{item.id}</TableCell>
                  <TableCell align="center">{item.edges?.edgesOfPatientrightsInsurance?.insurancecompany}</TableCell>
                  <TableCell align="center">{item.edges?.edgesOfPatientrightsMedicalrecordstaff?.name}</TableCell>

                  <TableCell align="center">{item.edges?.edgesOfPatientrightsAbilitypatientrights?.check}</TableCell>
                  <TableCell align="center">{item.edges?.edgesOfPatientrightsPatientrecord?.name}</TableCell>
                  <TableCell align="center">{item.responsible}</TableCell>
                  <TableCell align="center">{moment(item.permissionDate).format('DD/MM/YYYY HH.mm น.')}</TableCell>
                  <TableCell align="center">
                    <MoreInfo id={item}></MoreInfo>
                  </TableCell>




                </TableRow>
              ))}
          </TableBody>
        </Table>
      </TableContainer>
    );
  } else {
    return (
      <TableContainer component={Paper}>
        <Table className={classes.table} aria-label="simple table">
          <TableHead>
            <TableRow>
              <TableCell align="center">ไม่พบสิทธ์</TableCell>
            </TableRow>

          </TableHead>

        </Table>
      </TableContainer>
    );
  }

}




