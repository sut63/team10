import React, { FC } from 'react';
import MoreInfo from './BillInfo'
import { Link as RouterLink } from 'react-router-dom';

import { Cookies } from 'react-cookie/cjs';
import { useEffect, useState } from 'react';
import { Avatar,TextField,Paper,makeStyles,Button} from '@material-ui/core';
import { Autocomplete, Alert } from '@material-ui/lab';
import { Content, Header, Page, pageTheme, ContentHeader, Link, } from '@backstage/core';

import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';

import { DefaultApi } from '../../api/apis';
import { EntUser } from '../../api/models/EntUser';
import { EntBill } from '../../api/models/EntBill';
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
  const [bills, setBills] = useState<EntBill[]>(Array);
  const [Patientrecord, setPatientrecord] = React.useState<EntPatientrecord[]>([]);

  const getPatientrecord = async () => {
    const res = await http.listPatientrecord({ limit: 110, offset: 0 });
    setPatientrecord(res);
  };

  const getBills = async () => {
    const res = await http.getBill({ id: PatID })
    setLoading(false);
    setBills(res);
    if (res.length != 0) {
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

  }, [loading]);
  const SearchBill = async () => {
    getBills();
  };
  return (
    <Page theme={pageTheme.home}>
      <Header
        title={`ส่วนการเงิน`}>
        <Avatar alt="Remy Sharp" src={Users?.images as string} />
        <div style={{ marginLeft: 10 }}>{Name}</div>
      </Header>
      <Content>
        <ContentHeader title="ค้นหาใบเสร็จรับเงิน">
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
          &emsp;
          <Button
            onClick={() => {
              SearchBill();
            }}
            variant="contained"
            color="primary"
            style={{backgroundColor: "#00af91",padding: '6px 12px',}}
          >
            ค้นหา
               </Button>&emsp;
            <Link component={RouterLink} to="/createbill">
            <Button variant="contained">
              กลับ
           </Button>
          </Link>
        </ContentHeader>
        <TableContainer component={Paper}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="center">เลขที่ใบเสร็จ</TableCell>
                <TableCell align="center">ผู้ป่วย</TableCell>
                <TableCell align="center">ผู้รับเงิน</TableCell>
                <TableCell align="center">วันที่ชำระ</TableCell>
                <TableCell align="center">Manage</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {bills.map(bill => (
                  <TableRow>
                    <TableCell align="center">{bill.id}</TableCell>
                    <TableCell align="center">{bill.edges?.edgesOfUnpaybill?.edges?.edgesOfTreatment?.edges?.edgesOfPatientrecord?.name}</TableCell>
                    <TableCell align="center">{bill.edges?.edgesOfOfficer?.name}</TableCell>
                    <TableCell align="center">{bill.date}</TableCell>
                    <TableCell align="center">
                      <MoreInfo id={bill.id}></MoreInfo>
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