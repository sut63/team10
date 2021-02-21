import React, { FC } from 'react';
import MoreInfo from './prinfo'
import { Link as RouterLink } from 'react-router-dom';

import { Cookies } from 'react-cookie/cjs';
import { useEffect, useState } from 'react';
import { Avatar, TextField, Paper, makeStyles, Button } from '@material-ui/core';
import { Autocomplete, Alert } from '@material-ui/lab';
import { Content, Header, Page, pageTheme, ContentHeader, Link, } from '@backstage/core';

import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import AddBoxIcon from '@material-ui/icons/AddBox';

import { DefaultApi } from '../../api/apis';
import { EntUser } from '../../api/models/EntUser';
import { EntPatientrecord } from '../../api/models/EntPatientrecord';

const cookies = new Cookies();
const Name = cookies.get('Name');
const Img = cookies.get('Img');

const BillSearch: FC<{}> = () => {
  const http = new DefaultApi();
  const useStyles = makeStyles(theme => ({
    table: {
      minWidth: 650,
    },
    root: {
      '& > *': {
        borderBottom: 'unset',
      },
      width: '100%',
      maxWidth: 360,
      backgroundColor: theme.palette.background.paper,

    },
  }));

  const [PatID, setPatID] = React.useState<number>(0);
  const classes = useStyles();
  const [alert, setAlert] = React.useState(true);
  const [loading, setLoading] = React.useState(true);
  const [status, setStatus] = React.useState(false);
  const [Users, setUsers] = React.useState<Partial<EntUser>>();
  const [Patientrecords, setPatientrecords] = React.useState<EntPatientrecord[]>(Array);
  const [Patientrecord, setPatientrecord] = React.useState<EntPatientrecord[]>([]);

  const getPatientrecord = async () => {
    const res = await http.listPatientrecord({});
    setPatientrecord(res);
  };

  const StartPatientrecord = async () => {
    const res = await http.listPatientrecord({});
    setPatientrecords(res);
  };

  const getPatientrecords = async () => {
    const res = await http.getPatientrecord({ id: PatID })
    setLoading(false);
    setPatientrecords(res);
    if (res.length != 0 || PatID === undefined || PatID === null) {
      setStatus(true);
      setAlert(true);
    } else {
      setStatus(true);
      setAlert(false);
    }
    setTimeout(() => {
      setStatus(false);
    }, 1000);

  };

  const handleChange = (event: any, value: unknown) => {
    console.log(value)
    Patientrecord.map(item => {
      if (item.name === value) {
        setPatID(item.id as number);
      } else if (value === null) {
        setPatID(0);
      }
    })
  };

  const TexthandleChang = (event: React.ChangeEvent<{ value: unknown; }>) => {
    console.log(event.target.value as string)
    Patientrecord.map(item => {
      if (item.name === event.target.value as string) {
        setPatID(item.id as number);
      } else if (event.target.value as string === null) {
        setPatID(0);
      }
    })
  }
  useEffect(() => {
    const getImg = async () => {
      const res = await http.getUser({ id: Number(Img) });
      setLoading(false);
      setUsers(res);
    };
    getImg();
    getPatientrecord();
    StartPatientrecord();

  }, [loading]);
  const SearchPatientrecord = async () => {
    getPatientrecords();
  };
  return (
    <Page theme={pageTheme.home}>
      <Header
        title={`Patientrecord`}>
        <Avatar alt="Remy Sharp" src={Users?.images as string} />
        <div style={{ marginLeft: 10 }}>{Name}</div>
      </Header>
      <Content>
        <ContentHeader title="ค้นหาข้อมูลผู้ป่วย">
          <br />
          {status ? (
            <div>
              {alert ? (
                <Alert severity="success">
                  พบข้อมูลผู้ป่วย
                </Alert>
              ) : (
                  <Alert severity="warning" style={{ marginTop: 20 }}>
                    ไม่พบข้อมูลผู้ป่วย
                  </Alert>
                )}
            </div>
          ) : null}
        &emsp;
        <Autocomplete
            id="patientname"
            freeSolo
            options={Patientrecord.map(option => option.name)}
            onChange={handleChange}
            closeText='Close'
            renderInput={(params) => (
              <TextField {...params} label="ชื่อผู้ป่วย" margin="normal" variant="outlined" onChange={TexthandleChang} style={{ width: "100ch" }} />
            )}
          />
          <Button
            onClick={() => {
              SearchPatientrecord();
            }}
            style={{ marginLeft: 10 }}
            variant="contained"
            color="primary"
          >
            ค้นหา
               </Button>&emsp;
          <Link component={RouterLink} to="/">
            <Button variant="contained" color="primary" style={{ backgroundColor: "#d500f9" }}>
              Home
           </Button>
          </Link>&emsp;
          <Link component={RouterLink} to="/createPatientrecord">
            <Button variant="contained" color="primary" style={{ backgroundColor: "#9500ae" }} startIcon={<AddBoxIcon />} size="large">
              ลงทะเบียนผู้ป่วยนอก
            </Button>
          </Link>
        </ContentHeader>
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
                <TableCell align="center">Manage</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {Patientrecords.map(item => (
                <TableRow>
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
                  <TableCell align="center">{item.edges?.edgesOfMedicalrecordstaff?.name}</TableCell>
                  <TableCell align="center">
                    <MoreInfo id={item.id}></MoreInfo>
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Content>
    </Page>
  );

};

export default BillSearch;