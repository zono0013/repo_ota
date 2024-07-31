<img width="856" alt="スクリーンショット 2024-07-31 22 31 18" src="https://github.com/user-attachments/assets/e7080a8b-c4f3-4506-b46b-c0ff09146036">エンドポイント

- API実装
    - ユーザー
        - 作成
          - エンドポイント: `POST /users`
          - リクエストボディ:
          ```json
          {
            "name": "ユーザー名"
          }
          ```
          
          <img width="857" alt="スクリーンショット 2024-07-31 22 23 41" src="https://github.com/user-attachments/assets/d5eb711e-fd14-47fd-928e-c0f28afd4047">

        - 取得
          - エンドポイント: `GET /users/{id}`
          - パスパラメータ: `id `(ユーザーID)
          
          <img width="862" alt="スクリーンショット 2024-07-31 22 24 55" src="https://github.com/user-attachments/assets/feb4f386-4160-4620-adf1-cb8d2de46b29">

        - 更新
          - エンドポイント: `PUT /users/{id}`
          - パスパラメータ: `id` (ユーザーID)
          - リクエストボディ:
          
          ```json
          {
            "name": "新しいユーザー名"
          }
          ```
          
          <img width="867" alt="スクリーンショット 2024-07-31 22 25 58" src="https://github.com/user-attachments/assets/1b07607d-5721-4385-8dfd-b0f956a7f2fa">

          - 更新後の取得
            
          <img width="866" alt="スクリーンショット 2024-07-31 22 26 29" src="https://github.com/user-attachments/assets/0fec8685-f5dd-4f30-8183-2f47f191513d">




    - レポート詳細
        - 作成
          - エンドポイント: `POST /reports`
          - リクエストボディ:
            ```json{
              "user_id": 1,
              "title": "レポートのタイトル",
              "word_count": 500,
              "writing_style": "formal",
              "language": "japanese"
            }
            ```
            
          <img width="864" alt="スクリーンショット 2024-07-31 22 28 01" src="https://github.com/user-attachments/assets/1d54a2cb-3342-44e0-ac54-4c9e125b0097">

          <img width="861" alt="スクリーンショット 2024-07-31 22 28 47" src="https://github.com/user-attachments/assets/ed13b652-38ec-429c-aa01-32303264403e">

        - 一覧取得
          - エンドポイント: `GET /reports`
          <img width="863" alt="スクリーンショット 2024-07-31 22 29 19" src="https://github.com/user-attachments/assets/2a23400d-0d9c-43e9-8b52-e75515c4ad20">
          
        - 更新
          - エンドポイント: `PUT /reports/{id}`
          - パスパラメータ: `id` (レポートID)
          - リクエストボディ:
            ```json
            {
              "user_id": 1,
              "title": "更新されたレポートのタイトル",
              "word_count": 600,
              "writing_style": "informal",
              "language": "english"
            }
            ```
            
          <img width="852" alt="スクリーンショット 2024-07-31 22 31 51" src="https://github.com/user-attachments/assets/f64b52d2-19c9-4270-9ad5-112d4a97f073">

        - 単体取得
          - エンドポイント: `GET /reports/{id}`
          - パスパラメータ: `id` (レポートID)

          <img width="856" alt="スクリーンショット 2024-07-31 22 32 35" src="https://github.com/user-attachments/assets/64ff8894-371d-46c8-ab6a-03b12e41598f">

        - 削除
          - エンドポイント: `DELETE /reports/{id}`
          - パスパラメータ: `id` (レポートID)
            
          <img width="855" alt="スクリーンショット 2024-07-31 22 33 14" src="https://github.com/user-attachments/assets/da728add-4fd9-4ed8-8c7f-bd6a80731ae5">

          - 削除後の一覧取得
            
          <img width="859" alt="スクリーンショット 2024-07-31 22 33 41" src="https://github.com/user-attachments/assets/308eaf94-64af-45a3-835a-8de28d11fac3">

         

