import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';


import { DefaultApi } from '../../api/apis';
import { EntDoctorinfo } from '../../api/models/EntDoctorinfo';


import {
  Grid,
  Typography,
  Avatar,
  Button,
} from '@material-ui/core';


const useStyles = makeStyles(theme => ({
  table: {
    minWidth: 650,
  },
  formControl: {
    margin: theme.spacing(3),
    width: 200,
  },
  root: {
    width: '100%',
    maxWidth: 360,
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
  const [Doctorinfos, setDoctorinfos] = useState<EntDoctorinfo[]>(Array);
  const [loading, setLoading] = useState(true);
  const [nc, setnc] = useState(true);
  const [no, setno] = useState<EntDoctorinfo>({});

  const getDoctorinfos = async () => {
    const res = await api.listDoctorinfo();
    setLoading(false);
    setDoctorinfos(res);

  };

  useEffect(() => {

    getDoctorinfos();
    setnc(true);
  }, [loading]);



  const Show = async (i: any) => {

    setno(i);
    setnc(false);

  };
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

              <TableCell align="center">คำนำหน้า </TableCell>
              <TableCell align="center">ชื่อ </TableCell>
              <TableCell align="center">นามสกุล</TableCell>
              <TableCell align="center">เบอร์ติดต่อ</TableCell>
              <TableCell align="center">หมายเลขใบประกอบวิชาชีพ</TableCell>
              <TableCell align="center">ระดับการศึกษา</TableCell>
              <TableCell align="center">เเผนก</TableCell>
              <TableCell align="center">ห้องพัก</TableCell>

            </TableRow>
          </TableHead>
          <TableBody>
            {sim.sim === undefined
              ? null
              : sim.sim.map((item: any) => (
                <TableRow>

                  <TableCell align="center">{item.edges.edgesOfPrename.prefix}</TableCell>
                  <TableCell align="center">{item.doctorname}</TableCell>
                  <TableCell align="center">{item.doctorsurname}</TableCell>
                  <TableCell align="center">{item.telephonenumber}</TableCell>
                  <TableCell align="center">{item.licensenumber}</TableCell>
                  <TableCell align="center">{item.edges.edgesOfEducationlevel.level}</TableCell>
                  <TableCell align="center">{item.edges.edgesOfDepartment.department}</TableCell>
                  <TableCell align="center">{item.edges.edgesOfOfficeroom.roomnumber}</TableCell>




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
              <TableCell align="center">ไม่พบข้อมูล</TableCell>
            </TableRow>

          </TableHead>

        </Table>
      </TableContainer>
    );
  }
}


