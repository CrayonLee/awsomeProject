package main

import "fmt"

type Student struct{
	name string
	age int
	id int64
	class string
}

func NewStudent(name string,age int,id int64,class string ) *Student{
	return &Student{
		name:  name,
		age:   age,
		id:    id,
		class: class,
	}
}

type StuMgr struct{
	AllStudents []*Student
}

func NewStuMgr() *StuMgr{
	return &StuMgr{
		AllStudents: make([]*Student,0,100),
	}
}

func (s *StuMgr) AddStudnet (stu *Student)  {
	for _,v := range s.AllStudents {
		if v.name==stu.name{
			fmt.Printf("姓名为%s的学生已存在\n",stu.name)
			return
		}
	}
	s.AllStudents = append(s.AllStudents,stu)
}

func (s *StuMgr)EditStudent(stu *Student)  {
	for i,v := range s.AllStudents {
		if v.id == stu.id {
			s.AllStudents[i] = stu
			return
		}
	}
	fmt.Printf("姓名为%s的学生不存在\n", stu.name)
}

func (s *StuMgr) showAllStudnets()  {
	if(len(s.AllStudents)==0){
		fmt.Println("该系统中当前没有学生")
	}
	for _,v := range s.AllStudents {
		fmt.Printf("姓名:%s 年龄：%d id：%d 班级：%s",v.name,v.age,v.id,v.class)
	}

}

func (s *StuMgr) deleteStudent(id int64)  {
	for i,v := range s.AllStudents {
		if(v.id==id){
			//从切片中按照索引删除指定元素
			s.AllStudents=append(s.AllStudents[:i],s.AllStudents[i+1:]...)
		}

	}
}
