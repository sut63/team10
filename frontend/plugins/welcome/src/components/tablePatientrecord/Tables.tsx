import React, { useState, useEffect } from 'react';
import {  makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import { DefaultApi } from '../../api/apis';
import {
  Grid,
  Typography,
  Avatar,
  Button,
} from '@material-ui/core';
import { EntPatientrecord } from '../../api/models/EntPatientrecord';

const useStyles = makeStyles(theme => ({
  table: {
    minWidth: 650,
  },
  formControl: {
    margin: theme.spacing(3),
    width: 200,
  },
  root: {
    '& > *': {
      borderBottom: 'unset',
    },
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

export default function ComponanceTable(sim: any) {
  const classes = useStyles();
  const api = new DefaultApi();
  const [Patientrecord, setPatientrecord] = useState<EntPatientrecord[]>(Array);
  const [loading, setLoading] = useState(true);
  const [nc, setnc] = useState(true);
  const [no, setno] = useState<EntPatientrecord>({});

  const getPatientrecord = async () => {
    const res = await api.listPatientrecord({ limit: 100, offset: 0 });
    setLoading(false);
    setPatientrecord(res);
  };
  useEffect(() => {
    getPatientrecord();
    setnc(true);
  }, [loading]);

  const Show = async (i: any) => {
    setno(i);
    setnc(false);
  };
  var p = 0;

  for (var val of Patientrecord) {
    if (val.name === sim.sim || sim.sim === undefined || sim.sim === null) {
      p = p + 1
    }
  }
    if (p !== 0) {
      return (
        <TableContainer component={Paper}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
              <TableCell align="center">รหัสผู้ป่วย</TableCell>
              <TableCell align="center">คำนำหน้าชื่อ</TableCell>
              <TableCell align="center">ชื่อ-นามสกุล</TableCell>
              <TableCell align="center">เพศ</TableCell>
              <TableCell align="center">เลขบัตรประจำตัวประชาชน</TableCell>
              <TableCell align="center">อายุ</TableCell>
              <TableCell align="center">กรุ๊ปเลือด</TableCell>
              <TableCell align="center">โรคประจำตัว</TableCell>
              <TableCell align="center">แพ้ยา</TableCell>
              <TableCell align="center">เบอร์โทรที่ติดต่อได้</TableCell>
              <TableCell align="center">อีเมล์</TableCell>
              <TableCell align="center">ที่อยู่</TableCell>
              <TableCell align="center">วันเวลาที่ลงทะเบียนผู้ป่วยนอก</TableCell>
              <TableCell align="center">พนักงานเวชระเบียนที่ทำการลงบันทึก</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
            {Patientrecord === undefined
                ? null
                : Patientrecord.filter(i => i.name === sim.sim || sim.sim === undefined || sim.sim === null).map((item: any) => (
                <TableRow key={item.id}>
                <TableCell align="center">{item.id}</TableCell>
                <TableCell align="center">{item.edges?.edgesOfPrename?.prefix}</TableCell>
                <TableCell align="center">{item.name}</TableCell>
                <TableCell align="center">{item.edges?.edgesOfGender?.genderstatus}</TableCell>
                <TableCell align="center">{item.idcardnumber}</TableCell>
                <TableCell align="center">{item.age}</TableCell>
                <TableCell align="center">{item.edges?.edgesOfBloodtype?.bloodtype}</TableCell>
                <TableCell align="center">{item.disease}</TableCell>
                <TableCell align="center">{item.allergic}</TableCell>
                <TableCell align="center">{item.phonenumber}</TableCell>
                <TableCell align="center">{item.email}</TableCell>
                <TableCell align="center">{item.home}</TableCell>
                <TableCell align="center">{item.date}</TableCell>
                <TableCell align = "center">{item.edges?.edgesOfMedicalrecordstaff?.name}</TableCell>
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
                <TableCell align="center">ไม่พบข้อมูลผู้ป่วย</TableCell>
              </TableRow>
            </TableHead>
          </Table>
        </TableContainer>
      );
    } 
}