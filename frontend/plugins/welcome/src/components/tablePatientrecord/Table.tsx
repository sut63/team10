import React, { useState, useEffect } from 'react';
import { withStyles, createStyles, Theme, makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import { DefaultApi } from '../../api/apis';

import { EntPatientrecord } from '../../api/models/EntPatientrecord';

const StyledTableCell = withStyles((theme: Theme) =>
  createStyles({
    head: {
      backgroundColor: theme.palette.warning.main,
      color: theme.palette.common.white,
    },
  }),
)(TableCell);

const useStyles = makeStyles({
    table: {
      minWidth: 650,
    },
});

export default function ComponentsTable() {
    const classes = useStyles();
    const api = new DefaultApi();
    const [patientrecord, setPatientrecord] = React.useState<EntPatientrecord[]>([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const getPatientrecord = async () => {
            const res = await api.listPatientrecord({ offset: 0 });
            setLoading(false);
            setPatientrecord(res);
            console.log(res);
        };
        getPatientrecord();
    }, [loading]);

    return (
        <TableContainer component={Paper}>
              <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <StyledTableCell align="center">รหัสผู้ป่วย</StyledTableCell>
           <StyledTableCell align="center">คำนำหน้าชื่อ</StyledTableCell>
           <StyledTableCell align="center">ชื่อ-นามสกุล</StyledTableCell>
           <StyledTableCell align="center">เพศ</StyledTableCell>
           <StyledTableCell align="center">เลขบัตรประจำตัวประชาชน</StyledTableCell>
           <StyledTableCell align="center">อายุ</StyledTableCell>
           <StyledTableCell align="center">กรุ๊ปเลือด</StyledTableCell>
           <StyledTableCell align="center">โรคประจำตัว</StyledTableCell>
           <StyledTableCell align="center">แพ้ยา</StyledTableCell>
           <StyledTableCell align="center">เบอร์โทรที่ติดต่อได้</StyledTableCell>
           <StyledTableCell align="center">อีเมล์</StyledTableCell>
           <StyledTableCell align="center">ที่อยู่</StyledTableCell>
           <StyledTableCell align="center">วันเวลาที่ลงทะเบียนผู้ป่วยนอก</StyledTableCell>
           <StyledTableCell align="center">พนักงานเวชระเบียนที่ทำการลงบันทึก</StyledTableCell>
         </TableRow>
       </TableHead>
       <TableBody>
         {patientrecord.map(item => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.edges?.prename?.prefix}</TableCell>
             <TableCell align="center">{item.name}</TableCell>
             <TableCell align="center">{item.edges?.gender?.genderstatus}</TableCell>
             <TableCell align="center">{item.idcardnumber}</TableCell>
             <TableCell align="center">{item.age}</TableCell>
             <TableCell align="center">{item.bloodtype}</TableCell>
             <TableCell align="center">{item.disease}</TableCell>
             <TableCell align="center">{item.allergic}</TableCell>
             <TableCell align="center">{item.phonenumber}</TableCell>
             <TableCell align="center">{item.email}</TableCell>
             <TableCell align="center">{item.home}</TableCell>
             <TableCell align="center">{item.date}</TableCell>
            <TableCell align = "center">{item.edges?.medicalrecordstaff?.name}</TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
     </TableContainer>
    );
}