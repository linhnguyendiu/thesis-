package service

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/repository"
	"image"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/fogleman/gg"
	"github.com/go-redis/redis"
)

type CourseServiceImpl struct {
	repository.CourseRepository
	repository.OptionRepository
	repository.ExamResultRepository
	repository.ImageCourseRepository
	repository.ChapterRepository
	repository.LessonRepository
	repository.UserRepository
	repository.CertificateRepository
}

func (s *CourseServiceImpl) FindAllCourseIdByUserId(userId int) []web.CourseResponse {
	courses, err := s.CourseRepository.FindByUserId(userId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	coursesResponse := []web.CourseResponse{}
	for _, course := range courses {
		// countUsersEnrolled := s.CourseRepository.CountUsersEnrolled(course.Id)

		courseResponse := helper.ToCourseResponse(course)
		coursesResponse = append(coursesResponse, courseResponse)
	}

	return coursesResponse
}

func (s *CourseServiceImpl) CheckUserEnrollment(userId, courseId int) (bool, error) {
	return s.ExamResultRepository.HasUserEnrolledInCourse(userId, courseId)
}

func (s *CourseServiceImpl) FindByCategory(categoryName string) []web.CourseResponse {
	courses, err := s.CourseRepository.FindByCategory(categoryName)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	coursesResponse := []web.CourseResponse{}
	for _, course := range courses {
		// countUsersEnrolled := s.CourseRepository.CountUsersEnrolled(course.Id)

		courseResponse := helper.ToCourseResponse(course)
		coursesResponse = append(coursesResponse, courseResponse)
	}

	return coursesResponse
}

func (s *CourseServiceImpl) FindByUserId(userId int) []web.CourseResponse {
	courses, err := s.CourseRepository.FindByUserId(userId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	coursesResponse := []web.CourseResponse{}
	for _, course := range courses {
		// countUsersEnrolled := s.CourseRepository.CountUsersEnrolled(course.Id)

		courseResponse := helper.ToCourseResponse(course)
		coursesResponse = append(coursesResponse, courseResponse)
	}

	return coursesResponse
}

// func (s *CourseServiceImpl) UploadBanner(courseId int, pathFile string) bool {
// 	findById, err := s.CourseRepository.FindById(courseId)
// 	if err != nil {
// 		panic(helper.NewNotFoundError(err.Error()))
// 	}

// 	if findById.Banner != pathFile {
// 		if findById.Banner == "" {
// 			return updateWhenUploadBanner(findById, pathFile, s.CourseRepository)
// 		}
// 		os.Remove(findById.Banner)
// 		return updateWhenUploadBanner(findById, pathFile, s.CourseRepository)
// 	}

// 	return updateWhenUploadBanner(findById, pathFile, s.CourseRepository)
// }

func (s *CourseServiceImpl) UserEnrolled(userId int, courseId int) domain.UserCourse {
	_, err := s.CourseRepository.FindById(courseId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	userCourse := domain.UserCourse{
		CourseId: courseId,
		UserId:   userId,
	}

	usersEnrolled := s.CourseRepository.UsersEnrolled(userCourse)

	return usersEnrolled
}

func (s *CourseServiceImpl) FindAll() []web.CourseResponse {
	courses, err := s.CourseRepository.FindAll()
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	coursesResponse := []web.CourseResponse{}
	for _, course := range courses {
		// countUsersEnrolled := s.CourseRepository.CountUsersEnrolled(course.Id)

		courseResponse := helper.ToCourseResponse(course)
		coursesResponse = append(coursesResponse, courseResponse)
	}
	return coursesResponse
}

func (s *CourseServiceImpl) FindTop3Coures() []web.CourseResponse {
	courses, err := s.CourseRepository.FindTop3Course(3)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	coursesResponse := []web.CourseResponse{}
	for _, course := range courses {
		// countUsersEnrolled := s.CourseRepository.CountUsersEnrolled(course.Id)

		courseResponse := helper.ToCourseResponse(course)
		coursesResponse = append(coursesResponse, courseResponse)
	}
	return coursesResponse
}

func (s *CourseServiceImpl) FindByAuthorId(authorId int) []web.CourseResponse {
	courses, err := s.CourseRepository.FindByAuthorId(authorId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	coursesResponse := []web.CourseResponse{}
	for _, course := range courses {
		// countUsersEnrolled := s.CourseRepository.CountUsersEnrolled(course.Id)

		courseResponse := helper.ToCourseResponse(course)
		coursesResponse = append(coursesResponse, courseResponse)
	}
	return coursesResponse
}

func (s *CourseServiceImpl) FindByType(typeCourse string) []web.CourseResponse {
	findByType, err := s.CourseRepository.FindByType(typeCourse)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	coursesResponse := []web.CourseResponse{}
	for _, course := range findByType {
		// countUsersEnrolled := s.CourseRepository.CountUsersEnrolled(course.Id)

		courseResponse := helper.ToCourseResponse(course)
		coursesResponse = append(coursesResponse, courseResponse)
	}

	return coursesResponse

}

func (s *CourseServiceImpl) FindResultById(userId int, courseId int) web.ExamResultResponse {
	findById, err := s.ExamResultRepository.FindById(userId, courseId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	return helper.ToExamResultResponse(findById)
}

func (s *CourseServiceImpl) FindByTypeAndCategory(typeCourse string, cateName string) []web.CourseResponse {
	FindByTypeAndCategory, err := s.CourseRepository.FindByTypeAndCategory(typeCourse, cateName)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	coursesResponse := []web.CourseResponse{}
	for _, course := range FindByTypeAndCategory {
		// countUsersEnrolled := s.CourseRepository.CountUsersEnrolled(course.Id)

		courseResponse := helper.ToCourseResponse(course)
		coursesResponse = append(coursesResponse, courseResponse)
	}
	return coursesResponse

}

func (s *CourseServiceImpl) FindById(courseId int) web.CourseResponse {
	findById, err := s.CourseRepository.FindById(courseId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}
	//countUsersEnrolled := s.CourseRepository.CountUsersEnrolled(findById.Id)

	return helper.ToCourseResponse(findById)
	//return helper.ToCourseResponse(findById, countUsersEnrolled)
}

func (s *CourseServiceImpl) Create(request web.CourseCreateInput) web.CourseResponse {
	course := domain.Course{
		AuthorId:     request.AuthorId,
		Title:        request.Title,
		Type:         request.Type,
		Description:  request.Description,
		Price:        request.Price,
		Reward:       request.Reward,
		Category:     request.Category,
		DurationQuiz: request.DurationQuiz,
	}

	if course.AuthorId == 0 {
		panic(helper.NewUnauthorizedError("You're not an author"))
	}

	save := s.CourseRepository.Save(course)

	imageCourses := []domain.ImageCourse{}
	for _, imgInput := range request.ImageCourses {
		imageCourse := domain.ImageCourse{
			CourseId:    save.Id,
			ImageType:   imgInput.ImageType,
			Description: imgInput.Description,
			ImageAlt:    imgInput.ImageAlt,
		}
		savedImage := s.ImageCourseRepository.Save(imageCourse)
		imageCourses = append(imageCourses, savedImage)
	}
	course.ImageCourse = imageCourses

	chapters := []domain.Chapter{}
	for _, chapterInput := range request.Chapters {

		chapter := domain.Chapter{
			CourseId: save.Id,
			Title:    chapterInput.Title,
			InOrder:  chapterInput.InOrder,
		}
		savedChapter := s.ChapterRepository.Save(chapter)
		chapters = append(chapters, savedChapter)
		// Tạo các tùy chọn
		lessons := []domain.Lesson{}
		for _, lessonInput := range chapterInput.Lessons {
			lesson := domain.Lesson{
				ChapterId:    savedChapter.Id,
				Title:        lessonInput.Title,
				DurationTime: lessonInput.DurationTime,
				Description:  lessonInput.Description,
				Type:         lessonInput.Type,
				InOrder:      lessonInput.InOrder,
			}
			savedLesson := s.LessonRepository.Save(lesson)
			lessons = append(lessons, savedLesson)
		}
		savedChapter.Lesson = lessons
	}
	course.Chapter = chapters

	calculateTotalDuration, err := s.CourseRepository.CalculateTotalDuration(save.Id)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	countTotalLessonsInCourse, err := s.CourseRepository.CountTotalLessonsInCourse(save.Id)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}
	save.DurationToLearn = calculateTotalDuration
	save.LessonsCount = countTotalLessonsInCourse
	s.CourseRepository.Update(save)

	hash, err := helper.GenerateSHA256Hash(save)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	save.HashCourse = hash
	s.CourseRepository.Update(save)
	return helper.ToCourseResponse(save)
}

func (s *CourseServiceImpl) FindByKeyword(query string) ([]web.CourseResponse, error) {
	if query == "" {
		return nil, errors.New("user did not submit a valid query")
	}

	query = strings.ToLower(strings.TrimSpace(query))
	resultsMap := make(map[int]domain.Course)

	if !strings.Contains(query, " ") {
		courses, err := s.CourseRepository.FindByKeywords(query, 20)
		if err != nil {
			return nil, err
		}
		for _, course := range courses {
			resultsMap[course.Id] = course
		}
	} else {
		splitQuery := strings.Fields(query)
		for _, keyword := range splitQuery {
			courses, err := s.CourseRepository.FindByKeywords(keyword, 20)
			if err != nil {
				return nil, err
			}
			for _, course := range courses {
				resultsMap[course.Id] = course
			}
		}
	}

	coursesResponse := []web.CourseResponse{}
	for _, course := range resultsMap {

		courseResponse := helper.ToCourseResponse(course)
		coursesResponse = append(coursesResponse, courseResponse)
	}

	if len(coursesResponse) == 0 {
		return nil, errors.New("courses not found")
	}

	return coursesResponse, nil
}

func (s *CourseServiceImpl) EnrollCourse(input web.EnrollCourseInput) web.EnrollCourseResponse {
	examResult := domain.ExamResult{}
	examResult.CourseId = input.CourseId
	examResult.UserId = input.UserId
	examResult.EnrolledAt = input.EnrolledAt

	auth := helper.AuthGenerator(helper.Client)
	add, err := helper.Manage.BuyCourse(auth, big.NewInt(int64(input.UserId)), big.NewInt(int64(input.CourseId)))
	if err != nil {
		helper.PanicIfError(err)
	}
	log.Printf("buy course successfull", add)
	examResult.Status = true

	save := s.ExamResultRepository.Save(examResult)

	findById, err := s.CourseRepository.FindById(examResult.CourseId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	findById.UsersEnrolled += 1
	log.Printf("userenroll", findById.UsersEnrolled)
	saveCourse := s.CourseRepository.Update(findById)
	log.Printf("course", saveCourse)
	return helper.ToEnrollCourseResponse(save)
}

func (s *CourseServiceImpl) GetScore(ctx context.Context, request web.ExamRequest) web.ExamResultResponse {
	examResult, err := s.ExamResultRepository.FindById(request.UserId, request.CourseId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	isCompleted, err := s.IsCourseCompletedByUser(request.UserId, request.CourseId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}
	if !isCompleted {
		panic(helper.NewNotFoundError("User has not completed the course"))
	}

	findById, err := s.UserRepository.FindById(request.UserId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	course, err := s.CourseRepository.FindById(request.CourseId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	image, err := s.ImageCourseRepository.GetRandomImageByCourse(request.CourseId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	examResult.SubmittedAt = time.Now()
	log.Printf("time", time.Now())

	examResult.Score = 0

	totalQuestions, err := s.CourseRepository.GetTotalQuestionsByCourseId(examResult.CourseId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}
	examResult.TotalQuestions = int(totalQuestions)

	hash, err := helper.GenerateSHA256Hash(request)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	examResult.HashAnswer = hash

	for _, answerID := range request.Anwers {
		option, err := s.OptionRepository.FindById(answerID)
		if err != nil {
			panic(helper.NewNotFoundError(err.Error()))
		}
		if option.IsCorrect {
			examResult.Score++
		}
	}
	log.Printf("score", examResult.Score)

	examResult.Score = int((float64(examResult.Score) / float64(examResult.TotalQuestions)) * 10)
	log.Printf("score", examResult.Score)
	//save := s.ExamResultRepository.Save(examResult)

	// Kiểm tra lượt làm của người dùng
	userAttemptsKey := fmt.Sprintf("user:%d:course:%d:attempts", request.UserId, request.CourseId)
	attemptCount, err := helper.RedisCli.Get(ctx, userAttemptsKey).Int()
	if err == redis.Nil {
		attemptCount = 0
	}
	// } else if err != nil {
	// 	panic(err)
	// }

	attemptCount++

	if attemptCount == 1 {
		//Lần làm đầu tiên, lưu vào cơ sở dữ liệu
		auth := helper.AuthGenerator(helper.Client)
		add, err := helper.Manage.SubmitGrade(auth, big.NewInt(int64(examResult.UserId)), big.NewInt(int64(examResult.CourseId)), big.NewInt(int64(examResult.Score)), examResult.HashAnswer)
		if err != nil {
			helper.PanicIfError(err)
		}
		log.Printf("submit grade successfull", add)
		if examResult.Score > 7 {
			cert := domain.Certificate{
				UserName:   findById.LastName,
				CourseName: course.Title,
				Date:       examResult.SubmittedAt,
				CertType:   course.Type,
				ImageUri:   image.ImageAlt,
			}
			log.Printf("cert", cert)
			certificatePDF, err := GenerateCertPDF(cert)
			if err != nil {
				panic(helper.NewNotFoundError(err.Error()))
			}

			size := int64(len(certificatePDF))

			driveFileID, err := helper.CreateFile(cert.CourseName+cert.UserName+".png", size, certificatePDF)
			if err != nil {
				panic(helper.NewNotFoundError(err.Error()))
			}

			cert.CertUri = driveFileID
			certificateNFT := s.CertificateRepository.Save(cert)

			log.Printf("cert", certificateNFT)
			reward, err := helper.Manage.CheckAndTransferRewardCourse(auth, big.NewInt(int64(examResult.UserId)), big.NewInt(int64(examResult.CourseId)), big.NewInt(int64(certificateNFT.Id)), certificateNFT.ImageUri)
			if err != nil {
				helper.PanicIfError(err)
			}
			log.Printf("transac reward successfull", reward)
			txHash := reward.Hash().Hex()
			examResult.RewardAddress = txHash
			examResult.CertificateAddress = strconv.Itoa(certificateNFT.Id)
			findById.Balance = findById.Balance + course.Reward
		}

		save, err := s.ExamResultRepository.Update(examResult)
		if err != nil {
			panic(helper.NewNotFoundError(err.Error()))
		}

		if err := helper.RedisCli.Set(ctx, userAttemptsKey, 1, 0).Err(); err != nil {
			panic(err)
		}
		return helper.ToExamResultResponse(save)
	} else {
		// Lần làm thứ 2 trở đi, lưu vào Redis
		scoreKey := fmt.Sprintf("user:%d:course:%d:attempt:%d:score", request.UserId, request.CourseId, attemptCount)
		if err := helper.RedisCli.Set(ctx, scoreKey, examResult.Score, 0).Err(); err != nil {
			panic(err)
		}
		if err := helper.RedisCli.Set(ctx, userAttemptsKey, attemptCount, 0).Err(); err != nil {
			panic(err)
		}
		return web.ExamResultResponse{
			UserId:         request.UserId,
			CourseId:       request.CourseId,
			Attempt:        attemptCount,
			SubmittedAt:    examResult.SubmittedAt,
			Score:          examResult.Score,
			TotalQuestions: examResult.TotalQuestions,
		}
	}
}

func (s *CourseServiceImpl) IsCourseCompletedByUser(userId int, courseId int) (bool, error) {
	findById, err := s.CourseRepository.FindById(courseId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	countTotalLessonsInCourse, err := s.CourseRepository.CountTotalLessonsInCourse(findById.Id)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	completedLessons, err := s.CourseRepository.CountCompletedLessonsByUserInCourse(userId, findById.Id)
	if err != nil {
		return false, err
	}

	if int(completedLessons) >= countTotalLessonsInCourse {
		return true, nil
	}
	return false, nil
}

func GenerateCertPDF(cert domain.Certificate) ([]byte, error) {
	// Tạo khung ảnh với kích thước cố định
	const widthPx, heightPx = 1200, 800
	dc := gg.NewContext(widthPx, heightPx)

	// Background màu trắng
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// Vẽ tiêu đề chứng chỉ
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace("assets/font/BeVietnamPro-Black.ttf", 36); err != nil {
		return nil, err
	}
	dc.DrawStringAnchored("CHỨNG NHẬN HOÀN THÀNH KHÓA HỌC", widthPx/2, 100, 0.5, 0.5)

	// Vẽ hình ảnh chính
	response, err := http.Get(cert.ImageUri)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.New("received non 200 response code")
	}

	img, _, err := image.Decode(response.Body)
	if err != nil {
		return nil, err
	}
	// imgWidth, imgHeight := 300, 450
	dc.DrawImageAnchored(img, 200, 400, 0.5, 0.5)

	// Vẽ viền (border) cho phần text
	borderX := 550.0
	borderY := 150.0
	borderHeight := 500.0
	borderThickness := 2.0

	dc.SetLineWidth(borderThickness)
	dc.SetRGB(0, 0, 0)
	dc.DrawLine(borderX, borderY, borderX, borderY+borderHeight)
	dc.Stroke()

	// Vẽ thông tin chi tiết chứng chỉ
	if err := dc.LoadFontFace("assets/font/BeVietnamPro-Black.ttf", 24); err != nil {
		return nil, err
	}
	startX := 700.0
	startY := 200.0
	lineHeight := 40.0

	dc.DrawStringAnchored(cert.UserName, startX, startY, 0, 0.5)
	dc.DrawStringAnchored(cert.CourseName, startX, startY+lineHeight*2, 0, 0.5)
	dc.DrawStringAnchored(cert.Date.String(), startX, startY+4*lineHeight, 0, 0.5)
	dc.DrawStringAnchored(cert.CertType, startX, startY+6*lineHeight, 0, 0.5)

	var buf bytes.Buffer
	err = dc.EncodePNG(&buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func NewCourseService(courseRepository repository.CourseRepository, optionRepository repository.OptionRepository, examResultRepository repository.ExamResultRepository, imageCourseRepository repository.ImageCourseRepository, chapterRepository repository.ChapterRepository, lessonRepository repository.LessonRepository, userRepository repository.UserRepository, certificateRepository repository.CertificateRepository) CourseService {
	return &CourseServiceImpl{CourseRepository: courseRepository, OptionRepository: optionRepository, ExamResultRepository: examResultRepository, ImageCourseRepository: imageCourseRepository, ChapterRepository: chapterRepository, LessonRepository: lessonRepository, UserRepository: userRepository, CertificateRepository: certificateRepository}
}
