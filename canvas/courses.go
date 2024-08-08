package canvas

import "time"

type CanvasTerm struct {
	Id      int       `json:"id"`
	Name    string    `json:"name"`
	StartAt time.Time `json:"start_at"`
	EndAt   time.Time `json:"end_at,omitempty"`
}

type CanvasCourseProgress struct {
	RequirementCount          int       `json:"requirement_count"`
	RequirementCompletedCount int       `json:"requirement_completed_count"`
	NextRequirementURL        string    `json:"next_requirement_url,omitempty"`
	CompletedAt               time.Time `json:"completed_at,omitempty"`
}

/*
type CanvasCourse struct {
  // the unique identifier for the course
  id int // 370663,
  // the SIS identifier for the course, if defined. This field is only included if
  // the user has permission to view SIS information.
  sis_course_id int
  // the UUID of the course
  uuid string
  // the integration identifier for the course, if defined. This field is only
  // included if the user has permission to view SIS information.
  integration_id string
  // the unique identifier for the SIS import. This field is only included if the
  // user has permission to manage SIS information.
  sis_import_id int
  // the full name of the course. If the requesting user has set a nickname for
  // the course, the nickname will be shown here.
  name string
  // the course code
  course_code string
  // the actual course name. This field is returned only if the requesting user
  // has set a nickname for the course.
  "original_name": "InstructureCon-2012-01",
  // the current state of the course, also known as ‘status’.  The value will be
  // one of the following values: 'unpublished', 'available', 'completed', or
  // 'deleted'.  NOTE: When fetching a singular course that has a 'deleted'
  // workflow state value, an error will be returned with a message of 'The
  // specified resource does not exist.'
  "workflow_state": "available",
  // the account associated with the course
  "account_id": 81259,
  // the root account associated with the course
  "root_account_id": 81259,
  // the enrollment term associated with the course
  "enrollment_term_id": 34,
  // A list of grading periods associated with the course
  "grading_periods": null,
  // the grading standard associated with the course
  "grading_standard_id": 25,
  // the grade_passback_setting set on the course
  "grade_passback_setting": "nightly_sync",
  // the date the course was created.
  "created_at": "2012-05-01T00:00:00-06:00",
  // the start date for the course, if applicable
  "start_at": "2012-06-01T00:00:00-06:00",
  // the end date for the course, if applicable
  "end_at": "2012-09-01T00:00:00-06:00",
  // the course-set locale, if applicable
  "locale": "en",
  // A list of enrollments linking the current user to the course. for student
  // enrollments, grading information may be included if include[]=total_scores
  "enrollments": null,
  // optional: the total number of active and invited students in the course
  "total_students": 32,
  // course calendar
  "calendar": null,
  // the type of page that users will see when they first visit the course -
  // 'feed': Recent Activity Dashboard - 'wiki': Wiki Front Page - 'modules':
  // Course Modules/Sections Page - 'assignments': Course Assignments List -
  // 'syllabus': Course Syllabus Page other types may be added in the future
  "default_view": "feed",
  // optional: user-generated HTML for the course syllabus
  "syllabus_body": "<p>syllabus html goes here</p>",
  // optional: the number of submissions needing grading returned only if the
  // current user has grading rights and include[]=needs_grading_count
  "needs_grading_count": 17,
  // optional: the enrollment term object for the course returned only if
  // include[]=term
  "term": null,
  // optional: information on progress through the course returned only if
  // include[]=course_progress
  "course_progress": null,
  // weight final grade based on assignment group percentages
  "apply_assignment_group_weights": true,
  // optional: the permissions the user has for the course. returned only for a
  // single course and include[]=permissions
  "permissions": {"create_discussion_topic":true,"create_announcement":true},
  "is_public": true,
  "is_public_to_auth_users": true,
  "public_syllabus": true,
  "public_syllabus_to_auth": true,
  // optional: the public description of the course
  "public_description": "Come one, come all to InstructureCon 2012!",
  "storage_quota_mb": 5,
  "storage_quota_used_mb": 5,
  "hide_final_grades": false,
  "license": "Creative Commons",
  "allow_student_assignment_edits": false,
  "allow_wiki_comments": false,
  "allow_student_forum_attachments": false,
  "open_enrollment": true,
  "self_enrollment": false,
  "restrict_enrollments_to_course_dates": false,
  "course_format": "online",
  // optional: whether the course is set as a template (requires the Course
  // Templates feature)
  template bool
} */
