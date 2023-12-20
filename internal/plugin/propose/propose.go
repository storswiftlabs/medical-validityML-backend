package propose

import (
	"fmt"
	"medical-zkml-backend/internal/db"
	"medical-zkml-backend/internal/plugin/ipfs"
	"medical-zkml-backend/pkg/config"
	"os"
)

func init() {
	conf := config.NewConfig()
	db.InitMysql(conf)
}

func SavePropose() {
	type propose struct {
		disease string
		list    []string
	}

	proposes := make([]propose, 0, 7)

	proposes = append(proposes, propose{
		disease: "Acute Inflammations",
		list: []string{
			"1.Rest and Take Care of Yourself: When experiencing acute inflammation, it is important to get plenty of rest and allow your body to heal. Avoid overexertion and prioritize self-care.",
			"2.Apply Ice or Heat: Depending on the nature of the inflammation, applying ice packs or heat pads can help reduce pain and swelling. Ice is generally recommended for acute injuries, while heat may be more suitable for muscle stiffness or joint inflammation.",
			"3.Over-the-Counter Pain Relievers: Nonsteroidal anti-inflammatory drugs (NSAIDs), such as ibuprofen or naproxen, can help alleviate pain and reduce inflammation. Follow the recommended dosage guidelines and consult with a healthcare professional if necessary.",
			"4.Topical Treatments: Consider using topical creams or ointments containing ingredients like menthol, camphor, or capsaicin, which can provide temporary relief from localized inflammation and discomfort.",
			"5.Compression: For joint inflammation or swelling, consider using compression garments or bandages. These can help reduce swelling and provide support to the affected area.",
			"6.Elevate the Area: If possible, elevate the affected area to promote proper blood flow and reduce swelling.",
			"7.Maintain a Healthy Diet: Eat a balanced diet rich in fruits, vegetables, whole grains, and lean proteins. These foods provide essential nutrients to support the body's natural healing process and boost the immune system.",
			"8.Stay Hydrated: Drink an adequate amount of water throughout the day to keep your body hydrated. Proper hydration helps flush out toxins and supports overall health and recovery.",
			"9.Avoid Triggers: Identify and avoid any triggers that may exacerbate inflammation, such as certain foods, allergens, or environmental irritants.",
			"10.Consult a Healthcare Professional: If your symptoms worsen, persist, or if you have concerns about the severity of the inflammation, it is important to consult a healthcare professional for a proper diagnosis and appropriate treatment.",
		},
	})

	proposes = append(proposes, propose{
		disease: "Breast Cancer",
		list: []string{
			"1.Regular Breast Self-Exams: Perform monthly self-exams to check for any changes in your breasts, such as lumps, skin changes, or nipple discharge. Report any abnormalities to your healthcare provider.",
			"2.Clinical Breast Exams: Schedule regular clinical breast exams with your healthcare provider, who can detect any potential issues and discuss further steps if necessary.",
			"3.Mammograms: Starting at the recommended age (usually around 40 years old), undergo regular mammograms as advised by your doctor. Mammography can help detect breast cancer at an early stage, even before symptoms appear.",
			"4.Healthy Lifestyle: Maintain a healthy lifestyle by engaging in regular physical activity, adopting a balanced diet rich in fruits and vegetables, limiting alcohol consumption, and avoiding smoking.",
			"5.Know Your Family History: Understand your family's medical history, particularly regarding breast cancer. If you have a close relative (such as a mother or sister) who has had breast cancer, inform your doctor so they can assess your risk and recommend appropriate screening options.",
			"6.Breastfeeding: If you are able to and choose to breastfeed, research suggests it may reduce the risk of developing breast cancer.",
			"7.Early Detection and Prompt Action: If you notice any changes in your breasts or experience any breast-related concerns, seek medical attention promptly. Early detection and timely treatment can significantly improve outcomes.",
		},
	})

	proposes = append(proposes, propose{
		disease: "Chronic Kidney Disease",
		list: []string{
			"1.Control Blood Pressure: High blood pressure is a leading cause of chronic kidney disease. Work with your healthcare provider to monitor and control your blood pressure through lifestyle changes and/or medication.",
			"2.Control Blood Sugar: Diabetes is another significant cause of chronic kidney disease. If you have diabetes, maintain tight control of your blood sugar levels through diet, exercise, and medication.",
			"3.Quit Smoking: Smoking can worsen kidney damage and increase your risk of developing chronic kidney disease. Quitting smoking can help slow the progression of kidney disease and improve overall health.",
			"4.Watch Your Diet: A healthy diet can help manage chronic kidney disease. Limit sodium, phosphorus, and potassium intake, and eat a balanced diet rich in nutrients and low in unhealthy fats and processed foods.",
			"5.Stay Active: Regular exercise can help control blood pressure and improve overall health. Talk to your healthcare provider about appropriate exercise that fits your specific needs and abilities.",
			"6.Medication Management: Be sure to take medications as prescribed, and talk to your healthcare provider about any over-the-counter or herbal supplements that could affect kidney function.",
			"7.Regular Kidney Function Testing: Work with your healthcare provider to monitor your kidney function through regular blood and urine tests. Early detection can help slow the progression of chronic kidney disease and prevent further damage.",
		},
	})

	proposes = append(proposes, propose{
		disease: "Heart Disease",
		list: []string{
			"1.Healthy Diet: Follow a balanced diet that is low in saturated and trans fats, cholesterol, sodium, and added sugars. Include plenty of fruits, vegetables, whole grains, lean proteins, and healthy fats like those found in nuts and fish.",
			"2.Regular Physical Activity: Engage in regular aerobic exercise, such as brisk walking, jogging, cycling, or swimming. Aim for at least 150 minutes of moderate-intensity exercise or 75 minutes of vigorous exercise per week.",
			"3.Maintain a Healthy Weight: Achieve and maintain a healthy body weight by combining a nutritious diet with regular physical activity. This can help reduce the risk of heart disease.",
			"4.Avoid Tobacco Smoke: If you smoke, quit smoking. Avoid exposure to secondhand smoke as well, as it can increase the risk of heart disease.",
			"5.Control Blood Pressure: Monitor your blood pressure regularly and work with your healthcare provider to keep it within a healthy range. Lifestyle changes and medications may be necessary to manage high blood pressure effectively.",
			"6.Manage Cholesterol Levels: Keep your LDL (bad) cholesterol levels in check through a healthy diet, regular exercise, and medication if recommended by your healthcare provider.",
			"7.Control Blood Sugar: If you have diabetes, work with your healthcare provider to keep your blood sugar levels under control. High blood sugar can contribute to the development of heart disease.",
			"8.Limit Alcohol Intake: If you choose to drink alcohol, do so in moderation. For most adults, this means up to one drink per day for women and up to two drinks per day for men.",
			"9.Manage Stress: Find healthy ways to cope with stress, such as through relaxation techniques, exercise, hobbies, or talking to a therapist. Chronic stress can contribute to heart disease.",
		},
	})

	proposes = append(proposes, propose{
		disease: "Heart Failure Clinical Records",
		list: []string{
			"1.Medication Adherence: Take your prescribed medications as directed by your healthcare provider. Medications commonly prescribed for heart failure may include ACE inhibitors, beta blockers, diuretics, and others. These medications help manage symptoms and improve heart function.",
			"2.Monitor Fluid Intake: Keep track of your daily fluid intake, including water, beverages, and foods with high water content. In heart failure, excess fluid can accumulate in the body, leading to symptoms such as swelling and shortness of breath. Your healthcare provider may recommend limiting fluid intake and monitoring weight fluctuations.",
			"3.Low-sodium Diet: Follow a low-sodium diet to help reduce fluid retention. Limit your intake of processed foods, canned soups, fast food, and salty snacks. Instead, opt for fresh fruits, vegetables, lean proteins, and whole grains.",
			"4.Regular Exercise: Engage in a supervised exercise program as recommended by your healthcare provider. Exercise can help improve heart function, increase stamina, and reduce symptoms of heart failure. It's important to follow a program tailored to your specific needs and limitations.",
			"5.Weight Management: Maintain a healthy weight or work towards achieving a healthy weight through a combination of a balanced diet and regular exercise. Excess weight can strain the heart, exacerbating symptoms of heart failure.",
			"6.Quit Smoking: If you smoke, quit smoking. Smoking can worsen heart failure symptoms and increase the risk of complications. Talk to your healthcare provider about smoking cessation strategies.",
			"7.Limit Alcohol Intake: If you drink alcohol, do so in moderation or as advised by your healthcare provider. Excessive alcohol consumption can have a negative impact on heart health and worsen heart failure symptoms.",
			"8.Regular Medical Follow-up: Attend regular follow-up appointments with your healthcare provider. They will monitor your condition, make adjustments to your treatment plan if needed, and provide necessary guidance and support.",
			"9.Manage Stress: Find healthy ways to manage stress, as it can exacerbate heart failure symptoms. Engage in relaxation techniques, practice mindfulness, seek social support, or consider participating in stress management programs.",
			"10.Educate Yourself and Seek Support: Learn about heart failure, its symptoms, and how to manage the condition effectively. Seek support from healthcare professionals, support groups, or counseling services to help you cope with the challenges of living with heart failure.",
		},
	})

	proposes = append(proposes, propose{
		disease: "Lymphography",
		list: []string{
			"1.Consultation with a Healthcare Provider: If you need lymphatic imaging, schedule a consultation with a healthcare provider specializing in radiology or lymphatic imaging. They will assess your medical history and symptoms to determine the most appropriate imaging technique.",
			"2.Understand the Procedure: Familiarize yourself with the specific type of lymphatic imaging recommended for your case, such as lymphoscintigraphy, magnetic resonance lymphangiography (MRL), or computed tomography lymphangiography (CTLA). Research and discuss with your healthcare provider to understand the procedure, its benefits, and any potential risks or side effects.",
			"3.Follow Preparation Instructions: Follow any preparation instructions provided by your healthcare provider. This may include fasting for a certain period before the procedure, avoiding certain medications or substances, or wearing specific clothing.",
			"4.Communicate Any Allergies or Health Conditions: Inform your healthcare provider about any known allergies, previous reactions to contrast agents, or existing health conditions. This information is crucial for ensuring your safety during the imaging procedure.",
			"5.Ask Questions: Don't hesitate to ask questions or seek clarification about the procedure. Understanding the process and what to expect can help alleviate anxiety and ensure your cooperation throughout the imaging study.",
			"6.Take Medications as Directed: If your healthcare provider prescribes any medications or contrast agents for the imaging study, take them as directed. Follow the dosage and timing instructions carefully.",
			"7.Dress Comfortably: Wear loose, comfortable clothing on the day of the imaging procedure. This will allow easy access to the area being imaged and enhance your comfort during the process.",
			"8.Arrive on Time: Arrive at the imaging facility on time, allowing for additional paperwork or preliminary procedures if necessary. Being punctual helps ensure a smooth and efficient imaging experience.",
			"9.Communicate Any Discomfort: During the imaging procedure, communicate any discomfort or concerns to the technologist or healthcare provider. They can make adjustments or provide support to enhance your comfort.",
			"10.Follow Post-Imaging Instructions: After the lymphatic imaging, follow any post-imaging instructions provided by your healthcare provider. This may include avoiding strenuous activities, staying hydrated, or taking note of any unusual symptoms or side effects. Contact your healthcare provider if you have any concerns or questions after the procedure.",
		},
	})

	proposes = append(proposes, propose{
		disease: "Parkinsons",
		list: []string{
			"1.Consult with a Neurologist: If you suspect or have been diagnosed with Parkinson's disease, it is crucial to consult with a neurologist who specializes in movement disorders. They can provide an accurate diagnosis and guide your treatment plan.",
			"2.Medication Management: Work closely with your healthcare provider to develop a medication regimen that helps manage your symptoms effectively. Take your medications as prescribed and communicate any side effects or concerns to your healthcare team.",
			"3.Physical Therapy: Engage in regular physical therapy sessions to improve mobility, balance, and overall physical function. A physical therapist can design a personalized exercise program to address specific Parkinson's symptoms and help maintain muscle strength.",
			"4.Occupational Therapy: Consider working with an occupational therapist who can provide strategies and adaptations to enhance your daily activities and independence. They can suggest assistive devices to make tasks easier and safer.",
			"5.Speech Therapy: Parkinson's disease can affect speech and swallowing. Seek speech therapy to address speech difficulties and learn techniques to improve communication and swallowing function.",
			"6.Stay Active: Regular exercise is essential for maintaining overall health and managing Parkinson's symptoms. Engage in activities such as walking, swimming, biking, or yoga, as recommended by your healthcare provider.",
			"7.Balance and Fall Prevention: Take precautions to prevent falls by keeping your home environment safe and clutter-free. Use grab bars, non-slip mats, and proper lighting. Consider tai chi or other balance exercises to improve stability.",
			"8.Support Network: Join support groups or seek counseling to connect with others facing similar challenges. These resources can provide emotional support, information-sharing, and coping strategies.",
			"9.Healthy Lifestyle: Adopt a healthy lifestyle that includes a balanced diet, regular sleep patterns, stress management techniques, and avoiding smoking or excessive alcohol consumption.",
			"10.Stay Informed: Keep up-to-date with the latest research and developments in Parkinson's disease. Educate yourself and your loved ones about the condition to better understand its progression and treatment options.",
		},
	})

	for _, ele := range proposes {
		hash, err := ipfs.IPFS(ele.list)
		if err != nil {
			fmt.Println("Failed to import disease recommendations: ", err)
			os.Exit(8)
		}
		if err := db.RecordRecommendation(ele.disease, hash); err != nil {
			fmt.Println("Failed to import disease recommendations: ", err)
			os.Exit(8)
		}
	}
}
