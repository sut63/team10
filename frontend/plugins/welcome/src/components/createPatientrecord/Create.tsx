import React, { useEffect } from 'react';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import Select from '@material-ui/core/Select';
import { DefaultApi } from'../../api/apis';
import { Typography, Avatar } from '@material-ui/core'
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save';
import UndoIcon from '@material-ui/icons/Undo';
import { Alert } from '@material-ui/lab';
import { Link as RouterLink } from 'react-router-dom';
import TextField from '@material-ui/core/TextField';
import Swal from 'sweetalert2';

import { EntMedicalrecordstaff } from '../../api/models/EntMedicalrecordstaff';
import { EntBloodtype } from '../../api/models/EntBloodtype';
import { EntGender } from '../../api/models/EntGender';
import { EntPrename } from '../../api/models/EntPrename';
import { EntUser } from '../../api/models/EntUser';

import { Cookies } from 'react-cookie/cjs';//cookie

// header css
const HeaderCustom = {
  minHeight: '50px',
};

const cookies = new Cookies();
const MEDID = cookies.get('Med'); 
const Name = cookies.get('Name');
const Img = cookies.get('Img');


const useStyles = makeStyles((theme: Theme) =>
    createStyles({
      button: {
        display: 'block',
        marginTop: theme.spacing(2),
      },
      formControl: {
          width: 200,
        },
        selectEmpty: {
          marginTop: theme.spacing(2),
        },
        paper: {
          marginTop: theme.spacing(2),
          marginBottom: theme.spacing(2),
        },
        textField: {
            marginLeft: theme.spacing(1),
            marginRight: theme.spacing(1),
            width: 200,
        },
    }),
  );

  interface Patientrecord_Type {
    age?: string;
    allergic?: string;
    bloodtype?: number;
    date?: string;
    disease?: string;
    email?: string;
    gender?: number;
    home?: string;
    idcardnumber?: string;
    medicalrecordstaff?: number;
    name?: string;
    phonenumber?: string;
    prename?: number;
}

export  default  function Create() {
    
  const classes = useStyles();
  const api = new DefaultApi();

  const [idcardnumberError, setidcardnumberError] = React.useState('');
  const [phonenumberError, setphonenumberError] = React.useState('');
  const [emailError, setemailError] = React.useState('');

  const [prename, setPrename] = React.useState<EntPrename[]>([]);
  const [gender, setGender] = React.useState<EntGender[]>([]);
  const [bloodtype, setBloodtype] = React.useState<EntBloodtype[]>([]);
  const [medicalrecordstaff, setMedicalrecordstaff] = React.useState<Partial<EntMedicalrecordstaff>>();

  const [Patientrecord, setPatientrecord] = React.useState<Partial<Patientrecord_Type>>({});
   
  const [loading, setLoading] = React.useState(true);
  const [Users, setUsers] = React.useState<Partial<EntUser>>();
  
  // alert setting
  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 5000,
    timerProgressBar: true,
    didOpen: toast => {
      toast.addEventListener('mouseenter', Swal.stopTimer);
      toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
  });
  
  useEffect(() => {
    const getPrename = async () => {
        const res = await api.listPrename({ limit: 5, offset: 2 });
        setLoading(false);
        setPrename(res);
    };
    getPrename();
    const getGender = async () => {
        const res = await api.listGender({ limit: 2, offset: 0 });
        setLoading(false);
        setGender(res);
    };
    getGender();
    const getBloodtype = async () => {
      const res = await api.listBloodtype({ limit: 4, offset: 0 });
      setLoading(false);
      setBloodtype(res);
    };
    getBloodtype();
    const getMedicalrecordstaff = async () => {
        const res = await api.getMedicalrecordstaff({ id: Number(MEDID) });
        setLoading(false);
        setMedicalrecordstaff(res);
    };
    getMedicalrecordstaff();
    const getImg = async () => {
      const res = await api.getUser({ id: Number(Img) });
      setLoading(false);
      setUsers(res);
    };
    getImg();
    }, [loading]);
    
    //handle
    const handleChange = (

      event: React.ChangeEvent<{ name?: string; value: unknown }>,
    ) => {
      const name = event.target.name as keyof typeof CreatePatientrecord;
      const { value } = event.target;
      const validateValue = value as string
      checkPattern(name, validateValue)
      setPatientrecord({ ...Patientrecord, [name]: value,medicalrecordstaff:medicalrecordstaff?.id  });
    };

    // ฟังก์ชั่นสำหรับ validate idcardnumberError
  const validateidcardnumber = (val: string) => {
    return val.length == 13 ? true : false;
  }

  // ฟังก์ชั่นสำหรับ validate phonenumberError
  const validatephonenumber = (val: string) => {
    return val.length == 10 ? true : false;
  }

  // ฟังก์ชั่นสำหรับ validate emailError
  const validateemail = (email: string) => {
    const re = /^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    return re.test(email) ? true : false;
  }

  // สำหรับตรวจสอบรูปแบบข้อมูลที่กรอก ว่าเป็นไปตามที่กำหนดหรือไม่
  const checkPattern = (id: string, value: string) => {
    switch (id) {
      case 'Idcardnumber':
        validateidcardnumber(value) ? setidcardnumberError('') : setidcardnumberError('หมายเลยบัตรประชาชน 13 หลัก');
        return;
      case 'Phonenumber':
        validatephonenumber(value) ? setphonenumberError('') : setphonenumberError('หมายเลขโทรศัพท์ต้องเป็นตัวเลข 10 หลัก');
        return;
      case 'Email':
        validateemail(value) ? setemailError('') : setemailError('รูปแบบอีเมลไม่ถูกต้อง')
        return;
      default:
        return;
    }
  }
  
  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
  }
  const checkCaseSaveError = (field: string) => {
    switch(field) {
      case 'Idcardnumber':
        alertMessage("error","หมายเลยบัตรประชาชน 13 หลัก");
        return;
      case 'Phonenumber':
        alertMessage("error","หมายเลขโทรศัพท์ต้องเป็นตัวเลข 10 หลัก");
        return;
      case 'Email':
        alertMessage("error","รูปแบบอีเมลไม่ถูกต้อง");
        return;
      default:
        alertMessage("error","บันทึกข้อมูลไม่สำเร็จ");
        return;
    }
  }

  const CreatePatientrecord = async () => {
    const apiUrl = 'http://localhost:8080/api/v1/patientrecord';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(Patientrecord),
    };

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.status === true) {
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          checkCaseSaveError(data.error.Name);
        }
      });
  };

    return(
        <Page theme={pageTheme.home}>
          <Header style={HeaderCustom} title={`ลงทะเบียนผู้ป่วยนอก`}>
            <Avatar alt="Remy Sharp" src={Users?.images as string} />
            <div style={{ marginLeft: 10 }}>{Name}</div>
          </Header>
            <Content>
            <Typography variant="h6" gutterBottom align="center">
              <FormControl variant="outlined" className={classes.formControl}>
              <InputLabel>คำนำหน้าชื่อ</InputLabel>
              <Select
                name="prename"
                id="prename"
                label="คำนำหน้าชื่อ"
                value={Patientrecord.prename}
                onChange={handleChange}
              >
              {prename.map(item => {
                return (
                  <MenuItem value={item.id}>
                  {item.prefix}
                  </MenuItem>
                  );
                })}
              </Select>
              </FormControl> &emsp;

            <TextField 
            autoComplete="off"
            name="name"
            label="ชื่อ-นามสกุล" 
            variant="outlined" 
            type="string"
            value={Patientrecord.name}
            onChange={handleChange}
            /> &emsp;

              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เพศ</InputLabel>
                  <Select
                    name="gender"
                    id="gender"
                    label="เพศ"
                    value={Patientrecord.gender}
                    onChange={handleChange}
                  >
                  {gender.map(item => {
                    return (
                      <MenuItem value={item.id}>
                      {item.genderstatus}
                      </MenuItem>
                      );
                    })}
                  </Select>
              </FormControl>

            <div className={classes.paper}></div>
            <TextField 
            autoComplete="off"
            name="idcardnumber"
            label="เลขบัตรประจำตัวประชาชน" 
            variant="outlined" 
            type="string"
            value={Patientrecord.idcardnumber}
            onChange={handleChange}
            /> &emsp;

            <TextField 
            autoComplete="off"
            name="age"
            label="อายุ" 
            variant="outlined" 
            type="string"
            value={Patientrecord.age}
            onChange={handleChange}
            /> &emsp;
            
            <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>กรุ๊ปเลือด</InputLabel>
                  <Select
                    name="bloodtype"
                    id="bloodtype"
                    label="กรุ๊ปเลือด"
                    value={Patientrecord.bloodtype}
                    onChange={handleChange}
                  >
                  {bloodtype.map(item => {
                    return (
                      <MenuItem value={item.id}>
                      {item.bloodtype}
                      </MenuItem>
                      );
                    })}
                  </Select>
            </FormControl>
            
            <div className={classes.paper}></div>
            <TextField 
            autoComplete="off"
            name="disease"
            label="โรคประจำตัว" 
            variant="outlined" 
            type="string"
            style={{ width: "25ch"}}
            value={Patientrecord.disease}
            onChange={handleChange}
            /> &emsp;
            
            <TextField
            autoComplete="off"
            name="allergic"
            label="แพ้ยา"  
            variant="outlined" 
            type="string"
            style={{ width: "25ch"}}
            value={Patientrecord.allergic}
            onChange={handleChange}
            /> 
            
            <div className={classes.paper}></div>
            <TextField
            autoComplete="off" 
            name="phonenumber"
            label="เบอร์โทรที่ติดต่อได้" 
            variant="outlined"
            type="string"
            style={{ width: "25ch"}} 
            value={Patientrecord.phonenumber}
            onChange={handleChange}
            /> &emsp;
            
            <TextField 
            autoComplete="off"
            name="email"
            label="อีเมล์" 
            variant="outlined"
            type="string"
            style={{ width: "25ch"}} 
            value={Patientrecord.email}
            onChange={handleChange}
            /> 

            <div className={classes.paper}></div>
            <TextField
            autoComplete="off"
            name="home"
            label="ที่อยู่"
            multiline
            variant="outlined"
            type="string"
            style={{ width: "67ch"}}
            rows={3}
            value={Patientrecord.home}
            onChange={handleChange}
            />
              <br />
              <br /> พนักงานเวชระเบียน : {medicalrecordstaff?.name}
              <br />
              <br />
              <Typography variant="h6" gutterBottom  align="center">
              <Button
                onClick={() => {
                  CreatePatientrecord();
                }}
                variant="contained"
                color="primary"
                style={{backgroundColor: "#26c6da"}}
                startIcon={<SaveIcon />}
                >
                Submit
              </Button>
              <Button
                    style={{ marginLeft: 40 }}
                    startIcon={<UndoIcon />}
                    component={RouterLink}
                    to="/"
                    variant="contained"
                >
                  Back
                </Button>
                <Button
                    style={{ marginLeft: 40 }}
                    component={RouterLink}
                    to="/Patientrecord"
                    variant="contained"
                    color="secondary"
                >
                    SHOW
                </Button>
              </Typography>
            </Typography>
          </Content>
      </Page>
  );
}